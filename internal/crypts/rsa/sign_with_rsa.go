package rsa

import (
	"crypto"

	"github.com/starter-go/security/keys"
)

// SignWithRsaProvider 提供 RSA 签名算法
type SignWithRsaProvider struct {

	//starter:component
	_as func(keys.Registry) //starter:as(".")

	Enabled  bool //starter:inject("${security.algorithms.rsa.enabled}")
	Priority int  //starter:inject("${security.algorithms.rsa.priority}")
}

func (inst *SignWithRsaProvider) _impl() (keys.Registry, keys.Provider) {
	return inst, inst
}

// ListRegistrations ...
func (inst *SignWithRsaProvider) ListRegistrations() []*keys.Registration {

	b := &signWithRegBuilder{}
	b.provider = inst

	b.add(keys.PKCS1v15Padding, crypto.MD5)
	b.add(keys.PKCS1v15Padding, crypto.SHA1)
	b.add(keys.PKCS1v15Padding, crypto.SHA256)
	b.add(keys.PKCS1v15Padding, crypto.SHA512)

	b.add(keys.PSS, crypto.MD5)
	b.add(keys.PSS, crypto.SHA1)
	b.add(keys.PSS, crypto.SHA256)
	b.add(keys.PSS, crypto.SHA512)

	return b.items
}

// Name 返回 Provider Name
func (inst *SignWithRsaProvider) Name() string {
	return inst.PackageName()
}

// PackageName ...
func (inst *SignWithRsaProvider) PackageName() string {
	return "github.com/starter-go/security/internal/crypts/rsa#sign"
}

// Description ...
func (inst *SignWithRsaProvider) Description() string {
	return "Sign with RSA"
}

////////////////////////////////////////////////////////////////////////////////

type signWithRegBuilder struct {
	provider *SignWithRsaProvider
	items    []*keys.Registration
}

func (inst *signWithRegBuilder) add(padding keys.PaddingMode, hash crypto.Hash) {
	opt := &keys.SignatureOptions{
		KeyAlgorithm: "rsa",
		Hash:         hash,
		Padding:      padding,
	}
	algName := opt.Algorithm()
	alg := &signWithRsaAlgorithm{
		name:     algName,
		provider: inst.provider,
		opt:      *opt,
	}
	reg := &keys.Registration{
		Name:      algName.String(),
		Enabled:   inst.provider.Enabled,
		Priority:  inst.provider.Priority,
		Provider:  inst.provider,
		Algorithm: alg,
	}
	inst.items = append(inst.items, reg)
}

////////////////////////////////////////////////////////////////////////////////

type signWithRsaAlgorithm struct {
	name     keys.SignatureAlgorithmName
	provider *SignWithRsaProvider
	opt      keys.SignatureOptions
}

func (inst *signWithRsaAlgorithm) _impl() keys.SignatureAlgorithm {
	return inst
}

func (inst *signWithRsaAlgorithm) Name() string {
	return inst.name.String()
}

func (inst *signWithRsaAlgorithm) Provider() keys.Provider {
	return inst.provider
}

func (inst *signWithRsaAlgorithm) Options() *keys.SignatureOptions {
	dst := &keys.SignatureOptions{}
	*dst = inst.opt
	return dst
}

func (inst *signWithRsaAlgorithm) NewSigner(key keys.PrivateKey) keys.Signer {
	opt := inst.Options()
	return key.NewSigner(opt)
}

func (inst *signWithRsaAlgorithm) NewVerifier(key keys.PublicKey) keys.Verifier {
	opt := inst.Options()
	return key.NewVerifier(opt)
}
