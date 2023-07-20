package auth

// Authentication 表示一个身份验证请求
type Authentication interface {
	Mechanism() string
	Account() string
	Secret() []byte
}

// Authenticator 表示一个身份验证算法
type Authenticator interface {
	Authenticate(a Authentication) (User, error)
	Support(a Authentication) bool
}
