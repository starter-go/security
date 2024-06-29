package internal

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/jwt"
)

// SubjectServiceImpl ...
type SubjectServiceImpl struct {

	//starter:component

	_as func(rbac.SubjectService) //starter:as("#")

	JWTService jwt.Service //starter:inject("#")

}

func (inst *SubjectServiceImpl) _impl() rbac.SubjectService {
	return inst
}

// GetCurrent ...
func (inst *SubjectServiceImpl) GetCurrent(c context.Context) (*rbac.SubjectDTO, error) {

	dst := &rbac.SubjectDTO{}

	token, err := inst.JWTService.GetDTO(c)
	if err != nil {
		return dst, nil
	}

	session := token.Session
	user := session.Owner

	dst.Session = &session
	dst.ID = rbac.SubjectID(user)

	dst.Owner = session.Owner
	dst.NickName = session.User.NickName
	dst.Avatar = session.User.Avatar
	dst.Authenticated = session.Authenticated
	dst.Roles = session.User.Roles

	return dst, nil
}
