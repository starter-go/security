package auth

// Identity 表示一个经过验证的身份
type Identity interface {
	Class() string
	Name() string
	Mechanism() string
}

////////////////////////////////////////////////////////////////////////////////
// 扩展

// PhoneUserIdentity ...
type PhoneUserIdentity interface {
	Identity

	PhoneNumber() string
}

// EmailUserIdentity ...
type EmailUserIdentity interface {
	Identity

	EmailAddress() string
}

// AppUserIdentity ...
type AppUserIdentity interface {
	Identity

	UserName() string
}
