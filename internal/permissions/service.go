package permissions

import (
	"context"
	"fmt"

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

// List ...
func (inst *Service) List(c context.Context, q *rbac.PermissionQuery) ([]*rbac.PermissionDTO, error) {
	list, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return inst.Convertor.ConvertListE2D(c, list)
}

// Insert ...
func (inst *Service) Insert(c context.Context, o1 *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	o2, err := inst.Convertor.ConvertD2E(c, o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	return inst.Convertor.ConvertE2D(c, o3)
}

// Update ...
func (inst *Service) Update(c context.Context, id rbac.PermissionID, o *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Delete ...
func (inst *Service) Delete(c context.Context, id rbac.PermissionID) error {
	return inst.Dao.Delete(nil, id)
}
