package auth

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// User 表示经过验证的用户
type User interface {
	Name() rbac.UserName // 取用户名
	ID() rbac.UserID     // 取用户ID
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
