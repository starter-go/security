package rbac

import (
	"context"
)

// SessionDTO 表示会话信息
type SessionDTO struct {
	BaseDTO
}

// SessionService 是针对 SessionDTO 的服务
type SessionService interface {
	GetCurrent(c context.Context) (*SessionDTO, error)
}
