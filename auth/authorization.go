package auth

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// Authorization 表示一个授权请求
type Authorization interface {
	Context() context.Context
	User() rbac.User
}

// Authorizer 表示一个授权组件
type Authorizer interface {
	Authorize(a Authorization) error
	Support(a Authorization) bool
}

// Authorizers 表示一组授权组件
type Authorizers interface {
	Authorize(a Authorization) error
}
