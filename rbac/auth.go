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

	Action    string      `json:"action"`
	Mechanism string      `json:"mechanism"`
	Account   string      `json:"account"`
	Secret    lang.Base64 `json:"secret"`

	User        *UserDTO    `json:"user"`         // 用户信息 (optional)
	Step        int         `json:"step"`         // 表示验证的步骤
	NewPassword lang.Base64 `json:"new_password"` // 新的密码（用于注册，设置，重设密码）
	Success     bool        `json:"success"`      // 是否完成并且成功

	Properties map[string]string `json:"properties"` // 其它扩展属性
}

// AuthService 是针对 AuthDTO 的服务
type AuthService interface {
	Login(c context.Context, a *AuthDTO) (*AuthDTO, error)
}
