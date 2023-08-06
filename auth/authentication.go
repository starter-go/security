package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
)

// Authentication 表示一个身份验证请求
type Authentication interface {
	Context() context.Context
	Attributes() attributes.Table
	Parameters() parameters.Table
	Action() string
	Mechanism() string
	Account() string
	Secret() []byte
}

// Authenticator 表示一个身份验证算法
type Authenticator interface {

	// 验证用户身份
	Authenticate(a Authentication) (*Result, error)

	Support(a Authentication) bool
}
