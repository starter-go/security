package security

import "github.com/starter-go/application/properties"

// Token ... 表示一个安全令牌
type Token interface {
	GetProperties() properties.Table
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
