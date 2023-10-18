package rbac

import (
	"context"

	"github.com/starter-go/base/lang"
)

// 定义几种常用的授权动作
const (
	ActionLogin          = "login"
	ActionSignUp         = "sign-up"
	ActionResetPassword  = "reset-password"
	ActionChangePassword = "change-password"
	ActionSendCode       = "send-code"
)

// 定义几种常用的验证机制
const (
	MechanismPassword = "password"
	MechanismEmail    = "email"
	MechanismPhone    = "phone"
)

// AuthDTO 用于登录认证
type AuthDTO struct {
	BaseDTO

	Action    string      `json:"action"` // 如果为空，表示这条信息仅用于认证
	Mechanism string      `json:"mechanism"`
	Account   string      `json:"account"`
	Secret    lang.Base64 `json:"secret"`

	// User        *UserDTO    `json:"user"`         // 用户信息 (optional)
	// Step        int         `json:"step"`         // 表示验证的步骤
	// NewPassword lang.Base64 `json:"password"` // 新的密码（用于注册，设置，重设密码）
	// Success     bool        `json:"success"`      // 是否完成并且成功

	Parameters map[string]string `json:"parameters"` // 其它扩展属性
}

// AuthService 是针对 AuthDTO 的服务
type AuthService interface {
	Handle(c context.Context, a []*AuthDTO) error
}
