package keys

// Signature 包含签名信息
type Signature struct {
	Algorithm string
	Data      []byte
}

// Signer 提供签名方法
type Signer interface {

	// 生成签名
	Sign(want *Signature, key PrivateKey) (*Signature, error)
}

// Verifier 提供签名验证方法
type Verifier interface {

	// 验证签名
	Verify(signature *Signature, key PublicKey) error
}

// SignatureAlgorithm 表示签名算法
type SignatureAlgorithm interface {
	Algorithm

	Signer() Signer
	Verifier() Verifier
}
