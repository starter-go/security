package internal

import (
	"fmt"

	"github.com/starter-go/security/auth"
)

// AuthenticatorManagerImpl ...
type AuthenticatorManagerImpl struct {

	//starter:component
	_as func(auth.AuthenticatorManager) //starter:as("#")

	RegistryList []auth.Registry //starter:inject(".")

	cache []*auth.Registration
}

func (inst *AuthenticatorManagerImpl) _impl() auth.AuthenticatorManager {
	return inst
}

// FindFor ...
func (inst *AuthenticatorManagerImpl) FindFor(a auth.Authentication) (auth.Authenticator, error) {
	list := inst.getItems()
	for _, item := range list {
		if item.Mechanism.Support(a) {
			return item.Authenticator, nil
		}
	}
	return nil, fmt.Errorf("no Authenticator for Authentication, mechanism:%s", a.Mechanism())
}

func (inst *AuthenticatorManagerImpl) getItems() []*auth.Registration {
	list := inst.cache
	if list == nil {
		list = inst.loadItems()
		inst.cache = list
	}
	return list
}

func (inst *AuthenticatorManagerImpl) loadItems() []*auth.Registration {
	src := inst.RegistryList
	dst := make([]*auth.Registration, 0)
	for _, r1 := range src {
		list := r1.ListRegistrations()
		for _, r2 := range list {
			if inst.accept(r2) {
				dst = append(dst, r2)
			}
		}
	}
	return dst
}

func (inst *AuthenticatorManagerImpl) accept(r *auth.Registration) bool {
	return r.Enabled && (r.Mechanism != nil) && (r.Authenticator != nil)
}
