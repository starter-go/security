package users

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// Service ...
type Service struct {
	//starter:component
	_as func(rbac.UserService) //starter:as("#")

	Dao       rbac.UserDAO       //starter:inject("#")
	Convertor rbac.UserConvertor //starter:inject("#")
}

func (inst *Service) _impl() {
	inst._as(inst)
}

// Find ...
func (inst *Service) Find(c context.Context, id rbac.UserID) (*rbac.UserDTO, error) {
	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.Convertor.ConvertE2D(c, o1)
}
