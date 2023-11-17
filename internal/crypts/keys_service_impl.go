package crypts

import (
	"fmt"

	"github.com/starter-go/security/keys"
)

// ServiceImpl 实现 keys.Service
type ServiceImpl struct {

	//starter:component
	_as func(keys.Service) //starter:as("#")

	Manager keys.Manager //starter:inject("#")
}

func (inst *ServiceImpl) _impl() keys.Service {
	return inst
}

// GetManager ...
func (inst *ServiceImpl) GetManager() keys.Manager {
	return inst.Manager
}

// GetKeyPairAlgorithm ...
func (inst *ServiceImpl) GetKeyPairAlgorithm(algorithm string, selector func(reg *keys.Registration) bool) (keys.KeyPairAlgorithm, error) {
	alg, err := inst.Manager.Get(algorithm, selector)
	if err != nil {
		return nil, err
	}
	kpa, ok := alg.(keys.KeyPairAlgorithm)
	if !ok {
		return nil, fmt.Errorf("cannot cast algorithm[%s] as keys.KeyPairAlgorithm", algorithm)
	}
	return kpa, nil
}

// GetSecretKeyAlgorithm ...
func (inst *ServiceImpl) GetSecretKeyAlgorithm(algorithm string, selector func(reg *keys.Registration) bool) (keys.SecretKeyAlgorithm, error) {
	alg, err := inst.Manager.Get(algorithm, selector)
	if err != nil {
		return nil, err
	}
	ska, ok := alg.(keys.SecretKeyAlgorithm)
	if !ok {
		return nil, fmt.Errorf("cannot cast algorithm[%s] as keys.SecretKeyAlgorithm", algorithm)
	}
	return ska, nil
}

// GetSignatureAlgorithm ...
func (inst *ServiceImpl) GetSignatureAlgorithm(algorithm string, selector func(reg *keys.Registration) bool) (keys.SignatureAlgorithm, error) {
	alg, err := inst.Manager.Get(algorithm, selector)
	if err != nil {
		return nil, err
	}
	sa, ok := alg.(keys.SignatureAlgorithm)
	if !ok {
		return nil, fmt.Errorf("cannot cast algorithm[%s] as keys.SecretKeyAlgorithm", algorithm)
	}
	return sa, nil
}
