package keys

// PrivateKeyData 代表私钥 DTO
type PrivateKeyData struct {
	Algorithm   string
	Format      string
	ContentType string
	Content     []byte
}

// PrivateKey 代表私钥
type PrivateKey interface {
	Pair() KeyPair

	NewDecrypter(options *CryptOptions) Decrypter

	Export(want *PrivateKeyData) (*PrivateKeyData, error)
}

// PrivateKeyLoader 代表私钥
type PrivateKeyLoader interface {
	Load(o *PrivateKeyData) (PrivateKey, error)
}
