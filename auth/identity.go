package auth

// Identity 表示一个经过验证的身份
type Identity interface {
	Class() string
	Name() string
	Mechanism() string
}
