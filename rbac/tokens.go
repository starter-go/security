package rbac

import (
	"github.com/starter-go/base/lang"
)

// TokenDTO 表示令牌信息
type TokenDTO struct {
	BaseDTO

	MaxAge    lang.Milliseconds `json:"max_age"`    // 令牌的保质期
	ExpiredAt lang.Time         `json:"expired_at"` // 令牌的过期时间戳 (ExpiredAt = CreatedAt + MaxAge)
	Session   *SessionDTO       `json:"session"`    // 会话信息

	Properties map[string]string `json:"properties"`
}
