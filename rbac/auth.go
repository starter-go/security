package rbac

import (
	"context"
)

// AuthDTO 用于登录认证
type AuthDTO struct {
	BaseDTO

	Mechanism string
	Account   string
	Secret    []byte
}

// AuthService 是针对 AuthDTO 的服务
type AuthService interface {
	Login(c context.Context, a *AuthDTO) (*AuthDTO, error)
}
