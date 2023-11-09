package rbac

import "github.com/starter-go/base/lang"

// SessionDTO 表示会话信息
type SessionDTO struct {
	BaseDTO

	MaxAge        lang.Milliseconds `json:"max_age"`       // 会话的保质期
	ExpiredAt     lang.Time         `json:"expired_at"`    // 会话的过期时间戳 (ExpiredAt = CreatedAt + MaxAge)
	User          UserDTO           `json:"user"`          // 用户信息
	Authenticated bool              `json:"authenticated"` // 是否已验证

	Properties map[string]string `json:"properties"`
}
