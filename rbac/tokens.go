package rbac

import (
	"github.com/starter-go/base/lang"
)

// TokenDTO 表示令牌信息
type TokenDTO struct {
	BaseDTO

	ExpiredAt lang.Time   `json:"expired_at"` // 令牌的过期时间
	Session   *SessionDTO `json:"session"`    // 会话信息

	Properties map[string]string `json:"properties"`
}
