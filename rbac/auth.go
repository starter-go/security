package rbac

import (
	"context"

	"github.com/starter-go/base/lang"
)

// AuthDTO 用于登录认证
type AuthDTO struct {
	BaseDTO

	Mechanism string      `json:"mechanism"`
	Account   string      `json:"account"`
	Secret    lang.Base64 `json:"secret"`
}

// AuthService 是针对 AuthDTO 的服务
type AuthService interface {
	Login(c context.Context, a *AuthDTO) (*AuthDTO, error)
}
