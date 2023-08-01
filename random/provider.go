package random

// Provider 随机数提供方
type Provider interface {
	Source() Source
}

// ProviderRegistration 提供方注册信息
type ProviderRegistration struct {
	Name     string
	Enabled  bool
	Provider Provider
}

// ProviderRegistry 提供方注册接口
type ProviderRegistry interface {
	Registration() *ProviderRegistration
}
