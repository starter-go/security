package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/security/rbac"
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
	MechanismPhone    = rbac.MechanismPhone
)

// Request 包含 Authentication 和 Authorization 的公共方法
type Request interface {
	Context() context.Context
	Attributes() attributes.Table
	Parameters() parameters.Table
	Action() string
	Mechanism() string
}

// Mechanism  表示一个验证/授权机制组件
type Mechanism interface {
	Support(r Request) bool
}

// Service ...
type Service interface {
	// Authenticators
	// Authorizers
	// Login(c context.Context, a Authentication) error

	Handle(r Authentication) error
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

///////////////////////////////////

// // User 表示经过验证的用户
// type User interface {
// 	ID() rbac.UserID          // 用户ID
// 	Name() rbac.UserName      // 用户名
// 	Avatar() string           // 头像图片的 URL
// 	Roles() rbac.RoleNameList // 角色
// 	DisplayName() string      // 显示名称（昵称）
// }

// // Result 验证结果
// type Result struct {
// 	User      User
// 	Success   bool
// 	Challenge bool
// }

// // Authorizers 表示一组授权组件
// type Authorizers interface {
// 	Authorize(a Authorization) error
// }

// // Authenticators 表示一组身份验证算法
// type Authenticators interface {

// 	// 验证用户身份
// 	Authenticate(a Authentication) error
// }
