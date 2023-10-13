package rbac

import (
	"context"
)

// RoleID 是 Role 的实体 ID
type RoleID int64

// RoleName 是 Role 的正式名称
type RoleName string

// RoleNameList 是一组以逗号分隔的 RoleName
type RoleNameList string

// 定义一些常用的角色
const (
	RoleAdmin  RoleName = "admin"
	RoleAnonym RoleName = "anonym"
	RoleAny    RoleName = "any"
	RoleFriend RoleName = "friend"
	RoleGuest  RoleName = "guest"
	RoleOwner  RoleName = "owner"
	RoleRoot   RoleName = "root"
	RoleUser   RoleName = "user"
)

// RoleDTO 表示 Role 的 REST 网络对象
type RoleDTO struct {
	ID RoleID `json:"id"`

	BaseDTO

	Name RoleName `json:"name"`
}

// RoleQuery 查询参数
type RoleQuery struct {
	Pagination Pagination
}

// RoleService 是针对 RoleDTO 的服务
type RoleService interface {
	Insert(c context.Context, o *RoleDTO) (*RoleDTO, error)
	Update(c context.Context, id RoleID, o *RoleDTO) (*RoleDTO, error)
	Delete(c context.Context, id RoleID) error

	Find(c context.Context, id RoleID) (*RoleDTO, error)
	List(c context.Context, q *RoleQuery) ([]*RoleDTO, error)
}

////////////////////////////////////////////////////////////////////////////////

func (name RoleName) String() string {
	return string(name)
}
