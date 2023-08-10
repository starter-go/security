package rbac

import (
	"context"

	"github.com/starter-go/base/lang"
)

// SessionDTO 表示会话信息
type SessionDTO struct {
	BaseDTO

	ExpiredAt  lang.Time `json:"expired_at"` // 会话的过期时间
	User       UserDTO   `json:"user"`       // 用户信息
	Authorized bool      `json:"authorized"` // 是否已授权

	Properties map[string]string `json:"properties"`
}

// SessionService 是针对 SessionDTO 的服务
type SessionService interface {
	GetCurrent(c context.Context) (*SessionDTO, error)
}

// SessionProvider  会话的实现方案
type SessionProvider interface {
	Support(c context.Context) bool
	Current(c context.Context) (*SessionDTO, error)
}

// SessionRegistration 会话方案的注册信息
type SessionRegistration struct {
	Name     string
	Enabled  bool
	Priority int
	Provider SessionProvider
}

// SessionRegistry 会话方案的注册接口
type SessionRegistry interface {
	Registration() *SessionRegistration
}
