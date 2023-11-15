package keys

// Registration 表示密钥算法的注册信息
type Registration struct {
	Name      string
	Enabled   bool
	Priority  int
	Algorithm Algorithm
	Provider  Provider
}

// Registry 代表密钥算法的注册接口
type Registry interface {
	ListRegistrations() []*Registration
}
