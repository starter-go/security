package rbac

import (
	"context"
)

// PermissionID 是 Permission 的实体 ID
type PermissionID int64

// PermissionDTO 表示 Permission 的 REST 网络对象
type PermissionDTO struct {
	ID PermissionID `json:"id"`

	BaseDTO

	Method      string       `json:"method"`
	Path        string       `json:"path"`
	AcceptRoles RoleNameList `json:"accept_roles"`
	Enabled     bool         `json:"enabled"`
}

// PermissionQuery 查询参数
type PermissionQuery struct {
	Conditions Conditions
	Pagination Pagination
	All        bool // 查询全部条目
}

// PermissionService 是针对 PermissionDTO 的服务
type PermissionService interface {
	Insert(c context.Context, o *PermissionDTO) (*PermissionDTO, error)
	Update(c context.Context, id PermissionID, o *PermissionDTO) (*PermissionDTO, error)
	Delete(c context.Context, id PermissionID) error

	Find(c context.Context, id PermissionID) (*PermissionDTO, error)
	List(c context.Context, q *PermissionQuery) ([]*PermissionDTO, error)
	ListAll(c context.Context) ([]*PermissionDTO, error)

	GetCache() PermissionCache
}

// PermissionCache 是一个带缓存的 Permission 查询接口
type PermissionCache interface {
	Clear()
	Find(c context.Context, want *PermissionDTO) (*PermissionDTO, error)
}
