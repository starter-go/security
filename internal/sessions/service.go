package sessions

import (
	"context"

	"github.com/starter-go/security/jwt"
	"github.com/starter-go/security/rbac"
)

// Service ...
type Service struct {

	//starter:component
	_as func(rbac.SessionService) //starter:as("#")

	// Dao       rbac.SessionDAO       //starter:inject("#")
	// Convertor rbac.SessionConvertor //starter:inject("#")

	JWTser jwt.Service //starter:inject("#")
}

func (inst *Service) _impl() {
	inst._as(inst)
}

// GetCurrent ...
func (inst *Service) GetCurrent(c context.Context) (*rbac.SessionDTO, error) {
	session := &rbac.SessionDTO{}
	token, err := inst.JWTser.GetDTO(c)
	if err == nil && token != nil {
		*session = token.Session
	}
	return session, nil
}
