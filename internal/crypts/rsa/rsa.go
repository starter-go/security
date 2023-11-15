package rsa

import (
	"fmt"

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
	alg := &algorithm{myProvider: inst}
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

type algorithm struct {
	myProvider *Provider
}

func (inst *algorithm) _impl() (keys.Algorithm, keys.KeyPairAlgorithm) {
	return inst, inst
}

func (inst *algorithm) Name() string {
	return "rsa"
}

func (inst *algorithm) Provider() keys.Provider {
	return inst.myProvider
}

func (inst *algorithm) GetGenerator() keys.KeyPairGenerator {
	return &keypairGenerator{}
}

func (inst *algorithm) GetKeyPairLoader() keys.KeyPairLoader {
	return &keypairLoader{}
}

func (inst *algorithm) GetPublicKeyLoader() keys.PublicKeyLoader {
	return &publickeyLoader{}
}

func (inst *algorithm) GetPrivateKeyLoader() keys.PrivateKeyLoader {
	return &privatekeyLoader{}
}

////////////////////////////////////////////////////////////////////////////////

type publickeyLoader struct{}

func (inst *publickeyLoader) _impl() keys.PublicKeyLoader {
	return inst
}

func (inst *publickeyLoader) Load(o *keys.PublicKeyData) (keys.PublicKey, error) {
	return nil, fmt.Errorf("no impl ")
}

////////////////////////////////////////////////////////////////////////////////

type privatekeyLoader struct{}

func (inst *privatekeyLoader) _impl() keys.PrivateKeyLoader {
	return inst
}

func (inst *privatekeyLoader) Load(o *keys.PrivateKeyData) (keys.PrivateKey, error) {
	return nil, fmt.Errorf("no impl ")
}

////////////////////////////////////////////////////////////////////////////////

type keypairLoader struct{}

func (inst *keypairLoader) _impl() keys.KeyPairLoader {
	return inst
}

func (inst *keypairLoader) Load(o *keys.KeyPairData) (keys.KeyPair, error) {
	return nil, fmt.Errorf("no impl ")
}

////////////////////////////////////////////////////////////////////////////////

type keypairGenerator struct{}

func (inst *keypairGenerator) _impl() keys.KeyPairGenerator {
	return inst
}

func (inst *keypairGenerator) Generate(params *keys.KeyPairParams) (keys.KeyPair, error) {
	return nil, fmt.Errorf("no impl ")
}

////////////////////////////////////////////////////////////////////////////////
