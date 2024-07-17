package internal

import (
	"context"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security"
	"github.com/starter-go/security/subjects"
	"github.com/starter-go/vlog"
)

// use: RbacSubjectServiceImpl

type subjectCore struct {

	// context
	ctx            context.Context
	tokenService   rbac.TokenService
	sessionService rbac.SessionService

	// facade
	tokenFacade   security.Token
	sessionFacade security.Session

	// data
	tokenData   *rbac.TokenDTO
	sessionData *rbac.SessionDTO
}

func (inst *subjectCore) _impl() subjects.Subject { return inst }

func (inst *subjectCore) init() {
	inst.tokenFacade = &subjectTokenFacade{core: inst}
	inst.sessionFacade = &subjectSessionFacade{core: inst}
}

func (inst *subjectCore) GetContext() context.Context {
	return inst.ctx
}

func (inst *subjectCore) GetToken() security.Token {

	err := inst.fetchToken()
	if err != nil {
		vlog.Warn(err.Error())
		inst.tokenData = new(rbac.TokenDTO)
	}

	return inst.tokenFacade
}

func (inst *subjectCore) GetSession() security.Session {

	err := inst.fetchSession()
	if err != nil {
		vlog.Warn(err.Error())
		inst.sessionData = new(rbac.SessionDTO)
	}

	return inst.sessionFacade
}

func (inst *subjectCore) HasRole(role rbac.RoleName) bool {

	t := inst.GetToken()
	if t.HasRole(role) {
		return true
	}

	s := inst.GetSession()
	if s.HasRole(role) {
		return true
	}

	return false
}

func (inst *subjectCore) fetchToken() error {

	// 先检查是否已经缓存，如果没有才加载
	token := inst.tokenData
	if token != nil {
		return nil
	}

	// load
	ctx := inst.ctx
	token, err := inst.tokenService.GetCurrent(ctx)
	if err != nil {
		return err
	}

	inst.tokenData = token
	return nil
}

func (inst *subjectCore) fetchSession() error {

	// 先检查是否已经缓存，如果没有才加载
	ses := inst.sessionData
	if ses != nil {
		return nil
	}

	// get sid from token
	err := inst.fetchToken()
	if err != nil {
		return err
	}
	sid := inst.tokenData.Session

	// load
	ctx := inst.ctx
	ses, err = inst.sessionService.Find(ctx, sid)
	if err != nil {
		return err
	}

	inst.sessionData = ses
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type subjectSessionFacade struct {
	core  *subjectCore
	dirty bool
}

func (inst *subjectSessionFacade) _impl() security.Session { return inst }

func (inst *subjectSessionFacade) GetProperties() properties.Table {
	data := inst.core.sessionData
	src := data.Properties
	dst := properties.NewTable(nil)
	for k, v := range src {
		dst.SetProperty(k, v)
	}
	return dst
}

func (inst *subjectSessionFacade) SetProperty(name, value string) {
	data := inst.core.sessionData
	data.SetProperty(name, value)
	inst.dirty = true
}

func (inst *subjectSessionFacade) Get() *rbac.SessionDTO {
	data := inst.core.sessionData
	return data
}

func (inst *subjectSessionFacade) Set(s *rbac.SessionDTO) {
	if s == nil {
		return
	}
	inst.core.sessionData = s
	inst.dirty = true
}

func (inst *subjectSessionFacade) UserID() rbac.UserID {
	data := inst.core.sessionData
	return data.User
}

func (inst *subjectSessionFacade) SessionID() rbac.SessionID {
	data := inst.core.sessionData
	return data.ID
}

func (inst *subjectSessionFacade) UUID() lang.UUID {
	data := inst.core.sessionData
	return data.UUID
}

func (inst *subjectSessionFacade) UserName() rbac.UserName {
	data := inst.core.sessionData
	str := data.UUID.String()
	vlog.Warn("todo: no impl: subjectSessionFacade.UserName()")
	return rbac.UserName(str)
}

func (inst *subjectSessionFacade) Language() string {
	data := inst.core.sessionData
	return data.Language
}

func (inst *subjectSessionFacade) Nickname() string {
	data := inst.core.sessionData
	return data.Nickname
}

func (inst *subjectSessionFacade) Avatar() string {
	data := inst.core.sessionData
	return data.Avatar
}

func (inst *subjectSessionFacade) Roles() rbac.RoleNameList {
	data := inst.core.sessionData
	return data.Roles
}

func (inst *subjectSessionFacade) HasRole(role rbac.RoleName) bool {
	all := inst.Roles().List()
	for _, have := range all {
		if have == role {
			return true
		}
	}
	return false
}

func (inst *subjectSessionFacade) Authenticated() bool {
	data := inst.core.sessionData
	return data.Authenticated
}

func (inst *subjectSessionFacade) Create() error {

	// if !inst.dirty {
	// 	return nil
	// }

	ctx := inst.core.ctx
	ser := inst.core.sessionService
	item := inst.core.sessionData
	// id := item.ID

	item2, err := ser.Insert(ctx, item)
	if err != nil {
		return err
	}

	inst.core.sessionData = item2
	inst.dirty = false
	return nil
}

func (inst *subjectSessionFacade) Commit() error {

	if !inst.dirty {
		return nil
	}

	ctx := inst.core.ctx
	ser := inst.core.sessionService
	item := inst.core.sessionData
	id := item.ID

	item2, err := ser.Update(ctx, id, item)
	if err != nil {
		return err
	}

	inst.core.sessionData = item2
	inst.dirty = false
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type subjectTokenFacade struct {
	core  *subjectCore
	dirty bool
}

func (inst *subjectTokenFacade) _impl() security.Token { return inst }

func (inst *subjectTokenFacade) Get() *rbac.TokenDTO {
	data := inst.core.tokenData
	return data
}

func (inst *subjectTokenFacade) Set(t *rbac.TokenDTO) {
	if t == nil {
		return
	}
	inst.core.tokenData = t
	inst.dirty = true
}

func (inst *subjectTokenFacade) GetProperties() properties.Table {
	data := inst.core.tokenData
	src := data.Properties
	dst := properties.NewTable(nil)
	for k, v := range src {
		dst.SetProperty(k, v)
	}
	return dst
}

func (inst *subjectTokenFacade) SetProperty(name, value string) {
	data := inst.core.tokenData
	data.SetProperty(name, value)
	inst.dirty = true
}

func (inst *subjectTokenFacade) SetSessionID(sid rbac.SessionID) {
	data := inst.core.tokenData
	data.Session = sid
	inst.dirty = true
}

func (inst *subjectTokenFacade) GetSessionID() rbac.SessionID {
	data := inst.core.tokenData
	return data.Session
}

func (inst *subjectTokenFacade) Roles() rbac.RoleNameList {
	data := inst.core.tokenData
	return data.Roles
}

func (inst *subjectTokenFacade) HasRole(role rbac.RoleName) bool {
	all := inst.Roles().List()
	for _, have := range all {
		if have == role {
			return true
		}
	}
	return false
}

func (inst *subjectTokenFacade) Commit() error {

	if !inst.dirty {
		return nil
	}

	ctx := inst.core.ctx
	ser := inst.core.tokenService
	item := inst.core.tokenData

	item2, err := ser.PutCurrent(ctx, item)
	if err != nil {
		return err
	}

	inst.core.tokenData = item2
	inst.dirty = false
	return nil
}

////////////////////////////////////////////////////////////////////////////////
