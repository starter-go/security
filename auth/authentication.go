package auth

// Authentication 表示一个身份验证请求
type Authentication interface {
	Request

	Account() string
	Secret() []byte
}

// Authenticator 表示一个身份验证算法
type Authenticator interface {

	// 验证用户身份
	Authenticate(a Authentication) error
}
