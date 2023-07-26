package auth

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// User 表示经过验证的用户
type User interface {
	ID() rbac.UserID          // 用户ID
	Name() rbac.UserName      // 用户名
	Avatar() string           // 头像图片的 URL
	Roles() rbac.RoleNameList // 角色
	DisplayName() string      // 显示名称（昵称）
}

// Authorizers 表示一组授权组件
type Authorizers interface {
	Authorize(a Authorization) error
}

// Authenticators 表示一组身份验证算法
type Authenticators interface {
	Authenticate(a Authentication) (User, error)
}

// Service ...
type Service interface {
	Authenticators
	Authorizers

	Login(c context.Context, a Authentication) error
}

// Registration ... 注册信息
type Registration struct {
	Authenticator Authenticator
	Authorizer    Authorizer
	Priority      int
	Enabled       bool
}

// Registry ... 注册接口
type Registry interface {
	ListRegistrations() []*Registration
}
