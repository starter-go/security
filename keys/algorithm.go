package keys

// Algorithm 表示一种密码学算法
type Algorithm interface {
	Name() string
	Provider() Provider
}

// Provider 表示一个算法提供者
type Provider interface {
	Name() string
	PackageName() string
	Description() string
}
