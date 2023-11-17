package keys

import (
	"crypto"
	"strings"
)

// Signature 包含签名信息
type Signature struct {
	Algorithm SignatureAlgorithmName
	Digest    []byte
	Signature []byte
}

// Signer 提供签名方法
type Signer interface {

	// 取私钥
	PrivateKey() PrivateKey

	// 生成签名
	Sign(want *Signature) (*Signature, error)
}

// Verifier 提供签名验证方法
type Verifier interface {

	// 取公钥
	PublicKey() PublicKey

	// 验证签名
	Verify(signature *Signature) error
}

// SignatureAlgorithm 表示签名算法
type SignatureAlgorithm interface {
	Algorithm

	Options() *SignatureOptions

	NewSigner(key PrivateKey) Signer

	NewVerifier(key PublicKey) Verifier
}

////////////////////////////////////////////////////////////////////////////////

// SignatureOptions 表示签名选项
type SignatureOptions struct {
	KeyAlgorithm KeyAlgorithmName // 密钥算法
	Hash         crypto.Hash      // hash 算法 id
	Padding      PaddingMode      // 填充模式
}

// Algorithm 把选项组合转化为 SignatureAlgorithmName
func (inst *SignatureOptions) Algorithm() SignatureAlgorithmName {

	// [hash]with[key]/[padding]
	h := inst.Hash.String()
	k := inst.KeyAlgorithm.String()
	p := inst.Padding.String()
	builder := &strings.Builder{}

	h = inst.removeRune(h, '-')

	builder.WriteString(strings.ToUpper(h))
	builder.WriteString("with")
	builder.WriteString(strings.ToUpper(k))

	if p != "" {
		builder.WriteString("/")
		builder.WriteString(strings.ToUpper(p))
	}

	str := builder.String()
	return SignatureAlgorithmName(str)
}

func (inst *SignatureOptions) removeRune(str string, ch rune) string {
	if !strings.ContainsRune(str, ch) {
		return str
	}
	src := []rune(str)
	b := strings.Builder{}
	for _, ch2 := range src {
		if ch2 != ch {
			b.WriteRune(ch2)
		}
	}
	return b.String()
}

////////////////////////////////////////////////////////////////////////////////
