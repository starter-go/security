package permissions

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// Convertor ...
type Convertor struct {

	//starter:component
	_as func(rbac.PermissionConvertor) //starter:as("#")
}

func (inst *Convertor) _impl() rbac.PermissionConvertor {
	return inst
}

// ConvertE2D ...
func (inst *Convertor) ConvertE2D(c context.Context, src *rbac.PermissionEntity) (*rbac.PermissionDTO, error) {
	dst := &rbac.PermissionDTO{}
	dst.ID = src.ID
	rbac.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.AcceptRoles = src.AcceptRoles
	dst.Method = src.Method
	dst.Path = src.Path

	return dst, nil
}

// ConvertD2E ...
func (inst *Convertor) ConvertD2E(c context.Context, src *rbac.PermissionDTO) (*rbac.PermissionEntity, error) {
	dst := &rbac.PermissionEntity{}
	dst.ID = src.ID
	rbac.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.AcceptRoles = src.AcceptRoles
	dst.Method = src.Method
	dst.Path = src.Path

	return dst, nil
}

// ConvertListE2D ...
func (inst *Convertor) ConvertListE2D(c context.Context, src []*rbac.PermissionEntity) ([]*rbac.PermissionDTO, error) {
	dst := make([]*rbac.PermissionDTO, 0)
	for _, item1 := range src {
		item2, err := inst.ConvertE2D(c, item1)
		if err == nil && item2 != nil {
			dst = append(dst, item2)
		}
	}
	return dst, nil
}
