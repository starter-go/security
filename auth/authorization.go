package auth

// Authorization 表示一个授权请求
type Authorization interface {
	Request
}

// Authorizer 表示一个授权组件
type Authorizer interface {

	// 向用户授权
	Authorize(a Authorization) error
}
