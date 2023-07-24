package permissions

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// Service ...
type Service struct {
	//starter:component
	_as func(rbac.PermissionService) //starter:as("#")

	Dao       rbac.PermissionDAO       //starter:inject("#")
	Convertor rbac.PermissionConvertor //starter:inject("#")
}

func (inst *Service) _impl() {
	inst._as(inst)
}

// Find ...
func (inst *Service) Find(c context.Context, id rbac.PermissionID) (*rbac.PermissionDTO, error) {
	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.Convertor.ConvertE2D(c, o1)
}
