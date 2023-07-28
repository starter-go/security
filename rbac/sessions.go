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
