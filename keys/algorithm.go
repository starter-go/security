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

////////////////////////////////////////////////////////////////////////////////

// KeyAlgorithmName 表示密钥算法的名称
// (例如：'RSA' | 'AES' | 'ECDSA')
type KeyAlgorithmName string

func (name KeyAlgorithmName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////

// SignatureAlgorithmName 表示签名算法的名称
// (例如：'SHA384withECDSA' | 'SHA256withRSA/PSS')
type SignatureAlgorithmName string

func (name SignatureAlgorithmName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////
