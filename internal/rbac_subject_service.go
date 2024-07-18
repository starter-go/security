package internal

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/subjects"
)

// RbacSubjectServiceImpl ...
type RbacSubjectServiceImpl struct {

	//starter:component

	_as func(rbac.SubjectService) //starter:as("#")

	Tokens   rbac.TokenService   //starter:inject("#")
	Sessions rbac.SessionService //starter:inject("#")
}

func (inst *RbacSubjectServiceImpl) _impl() rbac.SubjectService { return inst }

// GetCurrent ...
func (inst *RbacSubjectServiceImpl) GetCurrent(c context.Context) (*rbac.SubjectDTO, error) {

	dst := &rbac.SubjectDTO{}
	sub, err := subjects.Current(c)
	if err != nil {
		dst.Token = &rbac.TokenDTO{}
		dst.Session = &rbac.SessionDTO{}
		return dst, nil
	}

	ses := sub.GetSession()
	tok := sub.GetToken()

	dst.Token = tok.Get()
	dst.Session = ses.Get()

	dst.Authenticated = dst.Session.Authenticated
	dst.CurrentUser = dst.Session.CurrentUser
	return dst, nil
}

func (inst *RbacSubjectServiceImpl) __older_GetCurrent(c context.Context) (*rbac.SubjectDTO, error) {

	sub := &rbac.SubjectDTO{}
	token1, err := inst.Tokens.GetCurrent(c)
	if err == nil {
		sub.Token = token1
		sesid := token1.Session
		session1, err := inst.Sessions.Find(c, sesid)
		if err == nil {
			sub.Session = session1
		} else {
			sub.Session = new(rbac.SessionDTO)
		}
	} else {
		sub.Token = new(rbac.TokenDTO)
		sub.Session = new(rbac.SessionDTO)
	}
	return sub, nil
}
