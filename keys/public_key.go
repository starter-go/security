package keys

// PublicKeyData 代表公钥的 DTO
type PublicKeyData struct {
	Algorithm   string
	Format      string
	ContentType string
	Content     []byte
}

// PublicKey 代表公钥
type PublicKey interface {
	NewEncrypter(options *CryptOptions) Encrypter

	NewVerifier(options *SignatureOptions) Verifier

	Export(want *PublicKeyData) (*PublicKeyData, error)
}

// PublicKeyLoader 代表公钥 loader
type PublicKeyLoader interface {
	Load(o *PublicKeyData) (PublicKey, error)
}
