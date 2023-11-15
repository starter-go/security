package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strings"

	"github.com/starter-go/security/keys"
	"github.com/starter-go/security/random"
)

// Provider 提供 AES 算法
type Provider struct {

	//starter:component
	_as func(keys.Registry) //starter:as(".")

	Rand random.Service //starter:inject("#")

	Enabled  bool //starter:inject("${security.algorithms.aes.enabled}")
	Priority int  //starter:inject("${security.algorithms.aes.priority}")
}

func (inst *Provider) _impl() (keys.Registry, keys.Provider) {
	return inst, inst
}

// ListRegistrations ...
func (inst *Provider) ListRegistrations() []*keys.Registration {
	ctx := &context{}
	alg := &algorithm{}

	ctx.algorithm = alg
	ctx.provider = inst
	alg.context = ctx

	r1 := &keys.Registration{
		Name:      alg.Name(),
		Enabled:   inst.Enabled,
		Priority:  inst.Priority,
		Algorithm: alg,
	}
	return []*keys.Registration{r1}
}

// Name 返回 Provider Name
func (inst *Provider) Name() string {
	return inst.PackageName()
}

// PackageName ...
func (inst *Provider) PackageName() string {
	return "github.com/starter-go/security/internal/crypts/aes"
}

// Description ...
func (inst *Provider) Description() string {
	return "-"
}

////////////////////////////////////////////////////////////////////////////////

type context struct {
	provider  *Provider
	algorithm *algorithm
	key       *secretkey
}

func (inst *context) clone() *context {
	dst := &context{
		provider:  inst.provider,
		algorithm: inst.algorithm,
		key:       inst.key,
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type algorithm struct {
	context *context
}

func (inst *algorithm) _impl() (keys.Algorithm, keys.SecretKeyAlgorithm) {
	return inst, inst
}

func (inst *algorithm) Name() string {
	return "aes"
}

func (inst *algorithm) Provider() keys.Provider {
	return inst.context.provider
}

func (inst *algorithm) GetGenerator() keys.SecretKeyGenerator {
	return &generator{context: inst.context}
}

func (inst *algorithm) GetLoader() keys.SecretKeyLoader {
	ctx := inst.context.clone()
	return &loader{context: ctx}
}

////////////////////////////////////////////////////////////////////////////////

type loader struct {
	context *context
}

func (inst *loader) Load(o *keys.SecretKeyData) (keys.SecretKey, error) {
	keydata, err := inst.loadKeyData(o)
	if err != nil {
		return nil, err
	}
	ctx := inst.context.clone()
	size := len(keydata) * 8
	skey := &secretkey{
		context: ctx,
		size:    size,
		secret:  keydata,
	}
	ctx.key = skey
	return skey, nil
}

func (inst *loader) trim(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return str
}

func (inst *loader) loadKeyData(o *keys.SecretKeyData) ([]byte, error) {

	if o == nil {
		return nil, fmt.Errorf("SecretKeyData is nil")
	}

	alg := inst.trim(o.Algorithm)
	ctype := inst.trim(o.ContentType)
	format := inst.trim(o.Format)
	data := o.Content
	size := len(data) * 8

	if alg == "aes" {
		if size == 128 || size == 192 || size == 256 {
			return data, nil
		}
	}

	return nil, fmt.Errorf("unsupported AES key format:%s size:%d content-type:%s", format, size, ctype)
}

////////////////////////////////////////////////////////////////////////////////

type generator struct {
	context *context
}

func (inst *generator) makeKey(key []byte) {
	r := inst.context.provider.Rand
	r.Reader().Read(key)
}

func (inst *generator) Generate(params *keys.SecretKeyParams) (keys.SecretKey, error) {
	const (
		aes128 = 128
		aes192 = 192
		aes256 = 256
	)
	size := params.Size
	switch size {
	case aes128:
	case aes192:
	case aes256:
		break
	default:
		return nil, fmt.Errorf("bad AES size: %d", size)
	}
	cb := size / 8
	key := make([]byte, cb)
	inst.makeKey(key)

	ctx := inst.context.clone()
	sk := &secretkey{context: ctx}
	sk.size = size
	sk.secret = key
	ctx.key = sk
	return sk, nil
}

////////////////////////////////////////////////////////////////////////////////

type secretkey struct {
	context *context
	size    int    // the key size
	secret  []byte // the key data

	block cipher.Block // 用于 get BlockSize
}

func (inst *secretkey) _impl() keys.SecretKey {
	return inst
}

func (inst *secretkey) Export(want *keys.SecretKeyData) (*keys.SecretKeyData, error) {
	src := inst.secret
	dst := make([]byte, len(src))
	copy(dst, src)
	have := &keys.SecretKeyData{}
	have.Algorithm = "aes"
	have.Format = "aes-raw-key"
	have.ContentType = "application/octet-stream"
	have.Content = dst
	return have, nil
}

func (inst *secretkey) NewEncrypter(options *keys.CryptOptions) keys.Encrypter {

	op2 := inst.context.key.prepareOptions(options)

	dst := &encrypter{}
	dst.context = inst.context.clone()
	dst.options = *op2
	return dst
}

func (inst *secretkey) NewDecrypter(options *keys.CryptOptions) keys.Decrypter {

	op2 := inst.context.key.prepareOptions(options)

	dst := &decrypter{}
	dst.context = inst.context.clone()
	dst.options = *op2
	return dst
}

// options : 靠后的选项会覆盖前之前的值
func (inst *secretkey) prepareOptions(options ...*keys.CryptOptions) *keys.CryptOptions {
	dst := &keys.CryptOptions{}
	for _, src := range options {
		if src == nil {
			continue
		}

		if src.Block != "" {
			dst.Block = src.Block
		}

		if src.Padding != "" {
			dst.Padding = src.Padding
		}

		if src.IV != nil {
			dst.IV = src.IV
		}
	}
	return dst
}

func (inst *secretkey) getBlock(cached cipher.Block) (cipher.Block, error) {
	if cached != nil {
		return cached, nil
	}
	keydata := inst.secret
	return aes.NewCipher(keydata)
}

func (inst *secretkey) BlockSize() int {
	b, err := inst.getBlock(inst.block)
	if err != nil {
		return 0
	}
	inst.block = b
	return b.BlockSize()
}

////////////////////////////////////////////////////////////////////////////////

type encrypter struct {
	context  *context
	options  keys.CryptOptions // the default options
	paddings keys.PaddingCache
	block    cipher.Block
}

func (inst *encrypter) _impl() keys.Encrypter {
	return inst
}

func (inst *encrypter) New(options *keys.CryptOptions) keys.Encrypter {

	op1 := &inst.options
	op2 := inst.context.key.prepareOptions(op1, options)

	dst := &encrypter{}
	dst.context = inst.context.clone()
	dst.options = *op2

	return dst
}

func (inst *encrypter) getBlock() (cipher.Block, error) {
	b, err := inst.context.key.getBlock(inst.block)
	inst.block = b
	return b, err
}

func (inst *encrypter) makeBlockMode(opt *keys.CryptOptions) (cipher.BlockMode, error) {
	iv := opt.IV
	mode := opt.Block
	block, err := inst.getBlock()
	if err != nil {
		return nil, err
	}
	switch mode {
	case keys.CBC:
		bm := cipher.NewCBCEncrypter(block, iv)
		return bm, nil
	}
	return nil, fmt.Errorf("unsupported block mode:%s", mode)
}

func (inst *encrypter) makeStreamMode(opt *keys.CryptOptions) (cipher.Stream, error) {
	iv := opt.IV
	mode := opt.Block
	block, err := inst.getBlock()
	if err != nil {
		return nil, err
	}
	switch mode {
	case keys.CFB:
		bm := cipher.NewCFBEncrypter(block, iv)
		return bm, nil

	case keys.CTR:
		bm := cipher.NewCTR(block, iv)
		return bm, nil
	}
	return nil, fmt.Errorf("unsupported stream mode:%s", mode)
}

func (inst *encrypter) Encrypt(src []byte, op1 *keys.CryptOptions) ([]byte, *keys.CryptOptions, error) {
	op2 := inst.context.key.prepareOptions(op1)
	// prepare crypt blocks
	bm, err := inst.makeBlockMode(op2)
	if err != nil {
		return nil, nil, err
	}
	//  padding
	padder, err := inst.paddings.Get(op2.Padding)
	if err != nil {
		return nil, nil, err
	}
	src, err = padder.Pad(src, bm.BlockSize())
	if err != nil {
		return nil, nil, err
	}
	// crypt
	dst := make([]byte, len(src))
	bm.CryptBlocks(dst, src)
	return dst, op2, nil
}

////////////////////////////////////////////////////////////////////////////////

type decrypter struct {
	context  *context
	options  keys.CryptOptions // the default options
	paddings keys.PaddingCache
	block    cipher.Block
}

func (inst *decrypter) _impl() keys.Decrypter {
	return inst
}

func (inst *decrypter) New(options *keys.CryptOptions) keys.Decrypter {
	op1 := &inst.options
	op2 := inst.context.key.prepareOptions(op1, options)

	dst := &decrypter{}
	dst.context = inst.context.clone()
	dst.options = *op2

	return dst
}

func (inst *decrypter) getBlock() (cipher.Block, error) {
	b, err := inst.context.key.getBlock(inst.block)
	inst.block = b
	return b, err
}

func (inst *decrypter) makeBlockMode(opt *keys.CryptOptions) (cipher.BlockMode, error) {
	iv := opt.IV
	mode := opt.Block
	block, err := inst.getBlock()
	if err != nil {
		return nil, err
	}
	switch mode {
	case keys.CBC:
		bm := cipher.NewCBCDecrypter(block, iv)
		return bm, nil
	}
	return nil, fmt.Errorf("unsupported block mode:%s", mode)
}

func (inst *decrypter) makeStreamMode(opt *keys.CryptOptions) (cipher.Stream, error) {
	iv := opt.IV
	mode := opt.Block
	block, err := inst.getBlock()
	if err != nil {
		return nil, err
	}
	switch mode {
	case keys.CFB:
		bm := cipher.NewCFBDecrypter(block, iv)
		return bm, nil

	case keys.CTR:
		bm := cipher.NewCTR(block, iv)
		return bm, nil
	}
	return nil, fmt.Errorf("unsupported stream mode:%s", mode)
}

func (inst *decrypter) Decrypt(src []byte, options *keys.CryptOptions) ([]byte, *keys.CryptOptions, error) {
	// crypt blocks
	op2 := inst.context.key.prepareOptions(options)
	dst := make([]byte, len(src))
	bm, err := inst.makeBlockMode(op2)
	if err != nil {
		return nil, nil, err
	}
	bm.CryptBlocks(dst, src)
	// un-padding
	padder, err := inst.paddings.Get(op2.Padding)
	if err != nil {
		return nil, nil, err
	}
	dst, err = padder.Unpad(dst, bm.BlockSize())
	if err != nil {
		return nil, nil, err
	}
	return dst, op2, nil
}

////////////////////////////////////////////////////////////////////////////////
