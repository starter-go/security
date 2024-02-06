package rsa

import (
	"crypto"
	"crypto/rand"
	crypotrand "crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"

	"github.com/starter-go/security/keys"
)

// Provider 提供 RSA 算法
type Provider struct {

	//starter:component
	_as func(keys.Registry) //starter:as(".")

	Enabled  bool //starter:inject("${security.algorithms.rsa.enabled}")
	Priority int  //starter:inject("${security.algorithms.rsa.priority}")
}

func (inst *Provider) _impl() (keys.Registry, keys.Provider) {
	return inst, inst
}

// ListRegistrations ...
func (inst *Provider) ListRegistrations() []*keys.Registration {

	ctx := &context{}
	alg := &algorithm{context: ctx}
	ctx.algorithm = alg
	ctx.provider = inst

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
	return "github.com/starter-go/security/internal/crypts/rsa"
}

// Description ...
func (inst *Provider) Description() string {
	return "-"
}

////////////////////////////////////////////////////////////////////////////////

type context struct {
	provider  *Provider
	algorithm *algorithm
	prikey    keys.PrivateKey
	pubkey    keys.PublicKey
}

func (inst *context) clone() *context {
	dst := &context{
		provider:  inst.provider,
		algorithm: inst.algorithm,
		pubkey:    inst.pubkey,
		prikey:    inst.prikey,
	}
	return dst
}

func (inst *context) computeSignatureOptions(src ...*keys.SignatureOptions) *keys.SignatureOptions {
	dst := &keys.SignatureOptions{}
	for _, item := range src {
		if item == nil {
			continue
		}
		if item.Hash > 0 {
			dst.Hash = item.Hash
		}
		if item.Padding != "" {
			dst.Padding = item.Padding
		}
	}
	return dst
}

func (inst *context) computeCryptOptions(src ...*keys.CryptOptions) *keys.CryptOptions {
	dst := &keys.CryptOptions{}
	for _, s := range src {
		if s == nil {
			continue
		}
		if s.Block != "" {
			dst.Block = s.Block
		}
		if s.Padding != "" {
			dst.Padding = s.Padding
		}
		if s.Hash != 0 {
			dst.Hash = s.Hash
		}
		if s.IV != nil {
			dst.IV = s.IV
		}
		if s.Label != nil {
			dst.Label = s.Label
		}
	}
	return dst
}

func (inst *context) getRandom() io.Reader {
	return crypotrand.Reader
}

////////////////////////////////////////////////////////////////////////////////

type algorithm struct {
	context *context
}

func (inst *algorithm) _impl() (keys.Algorithm, keys.KeyPairAlgorithm) {
	return inst, inst
}

func (inst *algorithm) Name() string {
	return "rsa"
}

func (inst *algorithm) Provider() keys.Provider {
	return inst.context.provider
}

func (inst *algorithm) GetGenerator() keys.KeyPairGenerator {
	return &keypairGenerator{context: inst.context}
}

func (inst *algorithm) GetKeyPairLoader() keys.KeyPairLoader {
	return &keypairLoader{context: inst.context}
}

func (inst *algorithm) GetPublicKeyLoader() keys.PublicKeyLoader {
	return &publickeyLoader{context: inst.context}
}

func (inst *algorithm) GetPrivateKeyLoader() keys.PrivateKeyLoader {
	return &privatekeyLoader{context: inst.context}
}

////////////////////////////////////////////////////////////////////////////////

type publickeyLoader struct {
	context *context
}

func (inst *publickeyLoader) _impl() keys.PublicKeyLoader {
	return inst
}

func (inst *publickeyLoader) Load(o *keys.PublicKeyData) (keys.PublicKey, error) {
	return nil, fmt.Errorf("no impl ")
}

////////////////////////////////////////////////////////////////////////////////

type privatekeyLoader struct {
	context *context
}

func (inst *privatekeyLoader) _impl() keys.PrivateKeyLoader {
	return inst
}

func (inst *privatekeyLoader) Load(o *keys.PrivateKeyData) (keys.PrivateKey, error) {
	return nil, fmt.Errorf("no impl ")
}

////////////////////////////////////////////////////////////////////////////////

type keypairLoader struct {
	context *context
}

func (inst *keypairLoader) _impl() keys.KeyPairLoader {
	return inst
}

func (inst *keypairLoader) Load(o *keys.KeyPairData) (keys.KeyPair, error) {
	return nil, fmt.Errorf("no impl ")
}

////////////////////////////////////////////////////////////////////////////////

type keypairGenerator struct {
	context *context
}

func (inst *keypairGenerator) _impl() keys.KeyPairGenerator {
	return inst
}

func (inst *keypairGenerator) Generate(params *keys.KeyPairParams) (keys.KeyPair, error) {

	if params == nil {
		params = &keys.KeyPairParams{Size: 1024 * 2}
	}

	key, err := rsa.GenerateKey(rand.Reader, params.Size)
	if err != nil {
		return nil, err
	}

	ctx := inst.context.clone()
	pubkey := &keyPublic{
		context: ctx,
		raw:     key.PublicKey,
	}
	prikey := &keyPrivate{
		context: ctx,
		raw:     *key,
	}
	pair := &keyPair{
		context: ctx,
	}
	ctx.prikey = prikey
	ctx.pubkey = pubkey
	return pair, nil
}

////////////////////////////////////////////////////////////////////////////////

type keyPrivate struct {
	context *context
	raw     rsa.PrivateKey
}

func (inst *keyPrivate) _impl() keys.PrivateKey {
	return inst
}

func (inst *keyPrivate) Native() crypto.PrivateKey {
	return inst.raw
}

func (inst *keyPrivate) NativeDecrypter() crypto.Decrypter {
	return &inst.raw
}

func (inst *keyPrivate) NativeSigner() crypto.Signer {
	return &inst.raw
}

func (inst *keyPrivate) Pair() keys.KeyPair {
	return &keyPair{
		context: inst.context,
	}
}

func (inst *keyPrivate) Export(want *keys.PrivateKeyData) (*keys.PrivateKeyData, error) {
	return nil, fmt.Errorf("no impl ")

}

func (inst *keyPrivate) NewDecrypter(options *keys.CryptOptions) keys.Decrypter {
	ctx := inst.context
	opt := ctx.computeCryptOptions(options)
	dst := &decrypter{
		context: ctx,
		options: opt,
		key:     inst,
	}
	return dst
}

func (inst *keyPrivate) NewSigner(options *keys.SignatureOptions) keys.Signer {
	ctx := inst.context
	opt := ctx.computeSignatureOptions(options)
	return &signer{key: inst, opt: opt}
}

////////////////////////////////////////////////////////////////////////////////

type keyPublic struct {
	context *context
	raw     rsa.PublicKey
}

func (inst *keyPublic) _impl() keys.PublicKey {
	return inst
}

func (inst *keyPublic) Native() crypto.PublicKey {
	return inst.raw
}

func (inst *keyPublic) Export(want *keys.PublicKeyData) (*keys.PublicKeyData, error) {
	return nil, fmt.Errorf("no impl ")

}

func (inst *keyPublic) NewEncrypter(options *keys.CryptOptions) keys.Encrypter {
	ctx := inst.context
	opt := ctx.computeCryptOptions(options)
	dst := &encrypter{
		context: ctx,
		options: opt,
		pubkey:  inst,
	}
	return dst
}

func (inst *keyPublic) NewVerifier(options *keys.SignatureOptions) keys.Verifier {
	ctx := inst.context
	opt := ctx.computeSignatureOptions(options)
	return &verifier{key: inst, opt: opt}
}

////////////////////////////////////////////////////////////////////////////////

type keyPair struct {
	context *context
}

func (inst *keyPair) _impl() keys.KeyPair {
	return inst
}

func (inst *keyPair) PublicKey() keys.PublicKey {
	return inst.context.pubkey
}

func (inst *keyPair) PrivateKey() keys.PrivateKey {
	return inst.context.prikey
}

func (inst *keyPair) Export(want *keys.KeyPairData) (*keys.KeyPairData, error) {
	return nil, fmt.Errorf("no impl ")

}

////////////////////////////////////////////////////////////////////////////////

type encrypter struct {
	context *context
	options *keys.CryptOptions
	pubkey  *keyPublic
}

func (inst *encrypter) _impl() keys.Encrypter {
	return inst
}

func (inst *encrypter) New(options *keys.CryptOptions) keys.Encrypter {
	ctx := inst.context
	opt := ctx.computeCryptOptions(inst.options, options)
	dst := &encrypter{
		context: ctx,
		options: opt,
	}
	return dst
}

func (inst *encrypter) Encrypt(data []byte, options *keys.CryptOptions) ([]byte, *keys.CryptOptions, error) {

	ctx := inst.context
	opt := ctx.computeCryptOptions(inst.options, options)
	key := &inst.pubkey.raw
	padding := opt.Padding
	random := inst.context.getRandom()

	if padding == keys.PKCS1v15Padding {
		mi, err := rsa.EncryptPKCS1v15(random, key, data)
		if err != nil {
			return nil, nil, err
		}
		return mi, opt, nil
	}

	if padding == keys.OAEP {
		h, err := keys.GetHash(opt.Hash)
		if err != nil {
			return nil, nil, err
		}
		label := opt.Label
		mi, err := rsa.EncryptOAEP(h, random, key, data, label)
		if err != nil {
			return nil, nil, err
		}
		return mi, opt, nil
	}

	return nil, nil, fmt.Errorf("unsupported RSA padding mode:%s", padding)
}

////////////////////////////////////////////////////////////////////////////////

type decrypter struct {
	context *context
	options *keys.CryptOptions
	key     *keyPrivate
}

func (inst *decrypter) _impl() keys.Decrypter {
	return inst
}

func (inst *decrypter) New(options *keys.CryptOptions) keys.Decrypter {
	ctx := inst.context
	opt := ctx.computeCryptOptions(inst.options, options)
	dst := &decrypter{
		context: ctx,
		options: opt,
	}
	return dst
}

func (inst *decrypter) Decrypt(data []byte, options *keys.CryptOptions) ([]byte, *keys.CryptOptions, error) {

	ctx := inst.context
	opt := ctx.computeCryptOptions(inst.options, options)
	padding := opt.Padding
	key := &inst.key.raw
	random := inst.context.getRandom()

	if padding == keys.PKCS1v15Padding {
		plain, err := rsa.DecryptPKCS1v15(random, key, data)
		if err != nil {
			return nil, nil, err
		}
		return plain, opt, nil
	}

	if padding == keys.OAEP {
		label := opt.Label
		h, err := keys.GetHash(opt.Hash)
		if err != nil {
			return nil, nil, err
		}
		plain, err := rsa.DecryptOAEP(h, random, key, data, label)
		if err != nil {
			return nil, nil, err
		}
		return plain, opt, nil
	}

	return nil, nil, fmt.Errorf("unsupported RSA padding mode:%s", padding)
}

////////////////////////////////////////////////////////////////////////////////

type signer struct {
	key *keyPrivate
	opt *keys.SignatureOptions
}

func (inst *signer) _impl() keys.Signer {
	return inst
}

func (inst *signer) PrivateKey() keys.PrivateKey {
	return inst.key
}

// 生成签名
func (inst *signer) Sign(want *keys.Signature) (*keys.Signature, error) {

	opt := inst.opt
	opt.KeyAlgorithm = "rsa"
	padding := opt.Padding
	random := inst.key.context.getRandom()
	pk := &inst.key.raw
	ha := opt.Hash
	dig := want.Digest

	dst := &keys.Signature{}
	dst.Algorithm = opt.Algorithm()
	dst.Digest = dig

	if padding == keys.PSS {
		sig, err := rsa.SignPSS(random, pk, ha, dig, nil)
		if err != nil {
			return nil, err
		}
		dst.Signature = sig
	}

	if padding == keys.PKCS1v15Padding {
		sig, err := rsa.SignPKCS1v15(random, pk, ha, dig)
		if err != nil {
			return nil, err
		}
		dst.Signature = sig
	}

	return dst, nil
}

////////////////////////////////////////////////////////////////////////////////

type verifier struct {
	key *keyPublic
	opt *keys.SignatureOptions
}

func (inst *verifier) _impl() keys.Verifier {
	return inst
}

func (inst *verifier) PublicKey() keys.PublicKey {
	return inst.key
}

// 验证签名
func (inst *verifier) Verify(signature *keys.Signature) error {

	padding := inst.opt.Padding
	pk := &inst.key.raw
	ha := inst.opt.Hash
	dig := signature.Digest
	sig := signature.Signature
	// random := inst.key.context.getRandom()

	if padding == keys.PSS {
		return rsa.VerifyPSS(pk, ha, dig, sig, nil)
	}

	if padding == keys.PKCS1v15Padding {
		return rsa.VerifyPKCS1v15(pk, ha, dig, sig)
	}

	return fmt.Errorf("bad padding mode for RSA signature, mode:%s", padding)
}

////////////////////////////////////////////////////////////////////////////////
