package keys

// PaddingMode 表示填充模式
type PaddingMode string

// 定义填充模式
const (
	NoPadding       PaddingMode = "NoPadding"
	PKCS5Padding    PaddingMode = "PKCS5Padding"
	PKCS7Padding    PaddingMode = "PKCS7Padding" // the default value
	PKCS1Padding    PaddingMode = "PKCS1Padding"
	PKCS1v15Padding PaddingMode = "PKCS1v15Padding"
	OAEP            PaddingMode = "OAEPadding" // optimal asymmetric encryption padding
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
	Hash    string // hash 算法名称
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
