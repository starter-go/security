package users

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// Convertor ...
type Convertor struct {

	//starter:component
	_as func(rbac.UserConvertor) //starter:as("#")
}

func (inst *Convertor) _impl() rbac.UserConvertor {
	return inst
}

// ConvertE2D ...
func (inst *Convertor) ConvertE2D(c context.Context, src *rbac.UserEntity) (*rbac.UserDTO, error) {
	dst := &rbac.UserDTO{}
	dst.ID = src.ID

	rbac.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Avatar = src.Avatar
	dst.Phone = src.Phone
	dst.Email = src.Email
	dst.NickName = src.Nickname
	dst.Name = src.Name
	dst.Roles = src.Roles

	return dst, nil
}

// ConvertD2E ...
func (inst *Convertor) ConvertD2E(c context.Context, src *rbac.UserDTO) (*rbac.UserEntity, error) {
	dst := &rbac.UserEntity{}
	dst.ID = src.ID

	rbac.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Avatar = src.Avatar
	dst.Phone = src.Phone
	dst.Email = src.Email
	dst.Nickname = src.NickName
	dst.Name = src.Name
	dst.Roles = src.Roles

	return dst, nil
}

// ConvertListE2D ...
func (inst *Convertor) ConvertListE2D(c context.Context, src []*rbac.UserEntity) ([]*rbac.UserDTO, error) {
	dst := make([]*rbac.UserDTO, 0)
	for _, item1 := range src {
		item2, err := inst.ConvertE2D(c, item1)
		if err == nil && item2 != nil {
			dst = append(dst, item2)
		}
	}
	return dst, nil
}
