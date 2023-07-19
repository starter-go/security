package auth

import "github.com/starter-go/security/rbac"

// Authentication 表示一个身份验证请求
type Authentication interface {
	Mechanism() string
	Account() string
	Secret() []byte
}

// Authenticator 表示一个身份验证算法
type Authenticator interface {
	Authenticate(a Authentication) (rbac.User, error)
	Support(a Authentication) bool
}

// Authenticators 表示一组身份验证算法
type Authenticators interface {
	Authenticate(a Authentication) (rbac.User, error)
}
