package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/rbac"
)

// 定义几种常用的授权动作
const (
	ActionLogin          = rbac.ActionLogin
	ActionSignUp         = rbac.ActionSignUp
	ActionResetPassword  = rbac.ActionResetPassword
	ActionChangePassword = rbac.ActionChangePassword
	ActionSendCode       = rbac.ActionSendCode
)

// 定义几种常用的验证机制
const (
	MechanismPassword = rbac.MechanismPassword
	MechanismEmail    = rbac.MechanismEmail
	MechanismPhone    = rbac.MechanismSMS
	MechanismSMS      = rbac.MechanismSMS
)

// Request 包含 Authentication 和 Authorization 的公共方法
type Request interface {
	Context() context.Context
	Attributes() attributes.Table
	Parameters() parameters.Table
	Feedback() Feedback

	Action() string
	Step() string
}

// Mechanism  表示一个验证/授权机制组件
type Mechanism interface {
	Support(r Request) bool
}

// Service ...
type Service interface {
	//  Authenticators
	// Authorizers
	// Login(c context.Context, a Authentication) error

	Authenticator

	Authorizer

	// 验证并授权
	Execute(r ...Request) error
}

// Registration ... 注册信息
type Registration struct {
	Authenticator Authenticator
	Authorizer    Authorizer
	Mechanism     Mechanism
	Priority      int
	Enabled       bool
}

// Registry ... 注册接口
type Registry interface {
	ListRegistrations() []*Registration
}
