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
	Resource    string // like 'method + ":" + path'
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

// PermissionService 是针对 PermissionDTO 的服务
type PermissionService interface {
	Find(c context.Context, id PermissionID) (*PermissionDTO, error)
}

// PermissionDAO 是数据库访问对象
type PermissionDAO interface {
	Find(db *gorm.DB, id PermissionID) (*PermissionEntity, error)
}
