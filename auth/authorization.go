package auth

import (
	"context"
)

// Authorization 表示一个授权请求
type Authorization interface {
	Context() context.Context
	User() User
}

// Authorizer 表示一个授权组件
type Authorizer interface {
	Authorize(a Authorization) error
	Support(a Authorization) bool
}

////////////////////////////////////////////////////////////////////////////////

type innerAuthorization struct {
	ctx  context.Context
	user User
}

func (inst *innerAuthorization) Context() context.Context {
	return inst.ctx
}

func (inst *innerAuthorization) User() User {
	return inst.user
}

// NewAuthorization 新建一个 Authorization 的实例
func NewAuthorization(c context.Context, user User) Authorization {
	return &innerAuthorization{
		ctx:  c,
		user: user,
	}
}

////////////////////////////////////////////////////////////////////////////////
