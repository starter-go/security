package keys

// SecretKeyData ...
type SecretKeyData struct {
	Algorithm   string
	Format      string
	ContentType string
	Content     []byte
}

// SecretKeyParams ...
type SecretKeyParams struct {
	Size int // key-size in bits
}

// SecretKey ...
type SecretKey interface {
	Export(want *SecretKeyData) (*SecretKeyData, error)

	NewEncrypter(options *CryptOptions) Encrypter

	NewDecrypter(options *CryptOptions) Decrypter

	BlockSize() int
}

// SecretKeyLoader ...
type SecretKeyLoader interface {
	Load(o *SecretKeyData) (SecretKey, error)
}

// SecretKeyGenerator ...
type SecretKeyGenerator interface {
	Generate(params *SecretKeyParams) (SecretKey, error)
}

// SecretKeyAlgorithm 表示一个对称密钥算法
type SecretKeyAlgorithm interface {
	Algorithm

	GetGenerator() SecretKeyGenerator

	GetLoader() SecretKeyLoader
}
