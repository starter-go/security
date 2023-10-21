package security

import (
	"context"
	"strings"

	"github.com/starter-go/base/context2"
	"github.com/starter-go/security/rbac"
)

const theSubjectBindingKey = "rbac.Subject#binding"

// Subject 代表当前主体
type Subject interface {
	// Init(ss SessionService)

	GetSession(create bool) Session

	HasRole(role rbac.RoleName) bool
}

// SetupSubject 把主体绑定到上下文中
func SetupSubject(c context.Context, ss SessionService) (Subject, error) {

	sub, err := GetSubject(c)
	if err == nil && sub != nil {
		return sub, nil
	}

	const key = theSubjectBindingKey
	kv, err := context2.GetValues(c)
	if err != nil {
		return nil, err
	}

	// make new subject
	s2 := &subject{context: c, sessionService: ss}
	sub = s2
	kv.SetValue(key, sub)
	return sub, nil
}

// GetSubject 从上下文中提取当前主体
func GetSubject(c context.Context) (Subject, error) {
	const key = theSubjectBindingKey
	kv, err := context2.GetValues(c)
	if err != nil {
		return nil, err
	}
	sub, ok := kv.GetValue(key).(Subject)
	if !ok {
		return sub, nil
	}
	return sub, nil
}

////////////////////////////////////////////////////////////////////////////////

type subject struct {
	context        context.Context
	session        Session
	sessionService SessionService
}

func (inst *subject) _impl(a Subject) {
	a = inst
}

func (inst *subject) Init(ss SessionService) {
	inst.sessionService = ss
}

func (inst *subject) GetSession(create bool) Session {
	s := inst.session
	if s == nil && create {
		ss := inst.sessionService
		ctx := inst.context
		if ss == nil {
			panic("use session without init subject")
		}
		s, err := ss.GetCurrent(ctx)
		if err != nil {
			panic(err)
		}
		inst.session = s
	}
	return s
}

func (inst *subject) HasRole(role rbac.RoleName) bool {
	session := inst.GetSession(true)
	roles := session.Roles()
	for _, have := range roles {
		str := strings.TrimSpace(have.String())
		if str == "" {
			continue
		}
		if have == role {
			return true
		}
	}
	return false
}
