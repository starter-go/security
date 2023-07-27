package rbac

import (
	"context"

	"gorm.io/gorm"
)

// PermissionID 是 Permission 的实体 ID
type PermissionID int64

// PermissionEntity 表示 Permission 的数据库实体
type PermissionEntity struct {
	ID PermissionID

	BaseEntity

	Method      string
	Path        string
	Resource    string `gorm:"unique"` // like 'method + ":" + path'
	AcceptRoles RoleNameList
}

// PermissionDTO 表示 Permission 的 REST 网络对象
type PermissionDTO struct {
	ID PermissionID `json:"id"`

	BaseDTO

	Method      string       `json:"method"`
	Path        string       `json:"path"`
	AcceptRoles RoleNameList `json:"accept_roles"`
}

// PermissionQuery 查询参数
type PermissionQuery struct {
	Pagination Pagination
}

// PermissionService 是针对 PermissionDTO 的服务
type PermissionService interface {
	Insert(c context.Context, o *PermissionDTO) (*PermissionDTO, error)
	Update(c context.Context, id PermissionID, o *PermissionDTO) (*PermissionDTO, error)
	Delete(c context.Context, id PermissionID) error

	Find(c context.Context, id PermissionID) (*PermissionDTO, error)
	List(c context.Context, q *PermissionQuery) ([]*PermissionDTO, error)
}

// PermissionDAO 是数据库访问对象
type PermissionDAO interface {
	Insert(db *gorm.DB, o *PermissionEntity) (*PermissionEntity, error)
	Update(db *gorm.DB, id PermissionID, updater func(*PermissionEntity)) (*PermissionEntity, error)
	Delete(db *gorm.DB, id PermissionID) error

	Find(db *gorm.DB, id PermissionID) (*PermissionEntity, error)
	List(db *gorm.DB, q *PermissionQuery) ([]*PermissionEntity, error)
}

// PermissionConvertor 负责 dto <==> entity 的转换
type PermissionConvertor interface {
	ConvertE2D(c context.Context, entity *PermissionEntity) (*PermissionDTO, error)
	ConvertD2E(c context.Context, dto *PermissionDTO) (*PermissionEntity, error)

	ConvertListE2D(c context.Context, entity []*PermissionEntity) ([]*PermissionDTO, error)
}
