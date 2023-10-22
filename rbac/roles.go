package rbac

import (
	"context"
	"sort"
	"strings"
)

// RoleID 是 Role 的实体 ID
type RoleID int64

// RoleName 是 Role 的正式名称
type RoleName string

// RoleNameList 是一组以逗号分隔的 RoleName
type RoleNameList string

// 定义一些常用的角色
const (
	RoleAdmin  RoleName = "admin"  // 管理员
	RoleAnonym RoleName = "anonym" // 匿名者
	RoleAny    RoleName = "any"    // 任何人
	RoleFriend RoleName = "friend" // 盆友
	RoleGuest  RoleName = "guest"  // 访客
	RoleOwner  RoleName = "owner"  // 资源持有者
	RoleRoot   RoleName = "root"   // 超级管理员
	RoleUser   RoleName = "user"   // 普通用户
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

////////////////////////////////////////////////////////////////////////////////

// List 拆分成单项列表
func (list RoleNameList) List() []RoleName {
	const sep = ","
	src := strings.Split(string(list), sep)
	dst := make([]RoleName, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		dst = append(dst, RoleName(item))
	}
	return dst
}

// Normalize 标准化
func (list RoleNameList) Normalize() RoleNameList {
	dst := make([]RoleName, 0)
	src := strings.Split(string(list), ",")
	strlist := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue // 排除空项
		}
		item = strings.ToLower(item)
		strlist = append(strlist, item)
	}
	sort.Strings(strlist)
	prev := ""
	for _, next := range strlist {
		if next == prev {
			continue // 排除重复项
		}
		prev = next
		dst = append(dst, RoleName(next))
	}
	return NewRoleNameList(dst...)
}

// NewRoleNameList 新建角色列表
func NewRoleNameList(names ...RoleName) RoleNameList {
	b := strings.Builder{}
	for i, name := range names {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(name.String())
	}
	str := b.String()
	return RoleNameList(str)
}

////////////////////////////////////////////////////////////////////////////////
