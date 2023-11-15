package keys

// KeyPairData ...
type KeyPairData struct {
	Algorithm   string
	Format      string
	ContentType string
	Content     []byte
}

// KeyPairParams ...
type KeyPairParams struct {
	Size int // key-size in bits
}

// KeyPair ...
type KeyPair interface {
	PublicKey() PublicKey
	PrivateKey() PrivateKey

	Export(want *KeyPairData) (*KeyPairData, error)
}

// KeyPairLoader ...
type KeyPairLoader interface {
	Load(o *KeyPairData) (KeyPair, error)
}

// KeyPairGenerator ...
type KeyPairGenerator interface {
	Generate(params *KeyPairParams) (KeyPair, error)
}

// KeyPairAlgorithm 表示一个非对称密钥对算法
type KeyPairAlgorithm interface {
	Algorithm

	GetGenerator() KeyPairGenerator

	GetKeyPairLoader() KeyPairLoader

	GetPublicKeyLoader() PublicKeyLoader

	GetPrivateKeyLoader() PrivateKeyLoader
}
