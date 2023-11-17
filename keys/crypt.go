package keys

import "crypto"

////////////////////////////////////////////////////////////////////////////////

// PaddingMode 表示填充模式
type PaddingMode string

func (m PaddingMode) String() string {
	return string(m)
}

////////////////////////////////////////////////////////////////////////////////

// 定义填充模式
const (
	NoPadding       PaddingMode = "No"
	PKCS5Padding    PaddingMode = "PKCS5"
	PKCS7Padding    PaddingMode = "PKCS7" // the default value
	PKCS1Padding    PaddingMode = "PKCS1"
	PKCS1v15Padding PaddingMode = "PKCS1v15"
	OAEP            PaddingMode = "OAEP" // optimal asymmetric encryption padding
	PSS             PaddingMode = "PSS"  // Probabilistic Signature Scheme
)

////////////////////////////////////////////////////////////////////////////////

// BlockMode 表示块模式
type BlockMode string

// 定义块模式
const (
	CBC BlockMode = "CBC" // Cipher Block Chaining
	EBC BlockMode = "EBC" // Electronic Codebook Book
	CTR BlockMode = "CTR" // Counter
	OCF BlockMode = "OCF" // Output FeedBack
	CFB BlockMode = "CFB" // Cipher FeedBack
)

////////////////////////////////////////////////////////////////////////////////

// CryptOptions 包含加密、解密的选项
type CryptOptions struct {
	Block   BlockMode
	Padding PaddingMode
	Hash    crypto.Hash // hash 算法 id
	Label   []byte
	IV      []byte // 初始化向量
}

////////////////////////////////////////////////////////////////////////////////

// Encrypter 提供加密计算
type Encrypter interface {
	New(options *CryptOptions) Encrypter

	Encrypt(data []byte, options *CryptOptions) ([]byte, *CryptOptions, error)
}

// Decrypter 提供解密计算
type Decrypter interface {
	New(options *CryptOptions) Decrypter

	Decrypt(data []byte, options *CryptOptions) ([]byte, *CryptOptions, error)
}
