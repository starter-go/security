package internal

import (
	"context"
	"fmt"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/subjects"
)

// DefaultSubjectLoader ...
type DefaultSubjectLoader struct {

	//starter:component

	_as func(subjects.Loader) //starter:as("#")

	Sessions rbac.SessionService //starter:inject("#")
	Tokens   rbac.TokenService   //starter:inject("#")

}

func (inst *DefaultSubjectLoader) _impl() subjects.Loader {
	return inst
}

// Load ...
func (inst *DefaultSubjectLoader) Load(c context.Context) (subjects.Subject, error) {

	if c == nil {
		return nil, fmt.Errorf("context is nil")
	}

	sub := &subjectCore{
		ctx:            c,
		sessionService: inst.Sessions,
		tokenService:   inst.Tokens,
	}
	sub.init()
	return sub, nil
}
