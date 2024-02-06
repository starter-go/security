package keys

import "crypto"

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

	NewSigner(options *SignatureOptions) Signer

	NativeSigner() crypto.Signer

	NativeDecrypter() crypto.Decrypter

	Native() crypto.PrivateKey

	Export(want *PrivateKeyData) (*PrivateKeyData, error)
}

// PrivateKeyLoader 代表私钥
type PrivateKeyLoader interface {
	Load(o *PrivateKeyData) (PrivateKey, error)
}
