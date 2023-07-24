package rbac

import (
	"context"

	"gorm.io/gorm"
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

// RoleEntity 表示 Role 的数据库实体
type RoleEntity struct {
	ID RoleID

	BaseEntity

	Name RoleName
}

// RoleDTO 表示 Role 的 REST 网络对象
type RoleDTO struct {
	ID RoleID `json:"id"`

	BaseDTO

	Name RoleName `json:"name"`
}

// RoleService 是针对 RoleDTO 的服务
type RoleService interface {
	Find(c context.Context, id RoleID) (*RoleDTO, error)
}

// RoleDAO 是数据库访问对象
type RoleDAO interface {
	Find(db *gorm.DB, id RoleID) (*RoleEntity, error)
}

// RoleConvertor 负责 dto <==> entity 的转换
type RoleConvertor interface {
	ConvertE2D(c context.Context, entity *RoleEntity) (*RoleDTO, error)
	ConvertD2E(c context.Context, dto *RoleDTO) (*RoleEntity, error)

	ConvertListE2D(c context.Context, entity []*RoleEntity) ([]*RoleDTO, error)
}
