package security

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac"
)

// Token ... 表示一个安全令牌
type Token interface {
	GetProperties() properties.Table

	SetProperty(name, value string)

	SetSessionID(sid rbac.SessionID)

	GetSessionID() rbac.SessionID

	Roles() rbac.RoleNameList

	HasRole(role rbac.RoleName) bool

	// 提交已修改的内容
	Commit() error
}

// TokenGetter ...
type TokenGetter interface {
	Get(ctx *Context) (Token, error)
}

// TokenSetter ...
type TokenSetter interface {
	Set(ctx *Context, t Token) error
}

// TokenLoader ...
type TokenLoader interface {
	Load(ctx *Context) (Token, error)
}

// TokenSaver ...
type TokenSaver interface {
	Save(ctx *Context, t Token) error
}

// TokenManager ...
type TokenManager interface {
	TokenGetter
	TokenSetter
	TokenLoader
	TokenSaver
}
