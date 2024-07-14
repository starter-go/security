package subjects

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security"
)

// Subject 表示当前的操作主体
type Subject interface {
	GetContext() context.Context

	GetToken() security.Token

	GetSession() security.Session

	HasRole(role rbac.RoleName) bool
}

// Loader 用于从上下文中加载 subject 对象
type Loader interface {
	Load(c context.Context) (Subject, error)
}
