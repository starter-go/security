package internal

import (
	"fmt"

	"github.com/starter-go/security/auth"
)

// AuthorizerManagerImpl ...
type AuthorizerManagerImpl struct {

	//starter:component
	_as func(auth.AuthorizerManager) //starter:as("#")

	RegistryList []auth.Registry //starter:inject(".")

	cache []*auth.Registration
}

func (inst *AuthorizerManagerImpl) _impl() auth.AuthorizerManager {
	return inst
}

// FindFor ...
func (inst *AuthorizerManagerImpl) FindFor(a auth.Authorization) (auth.Authorizer, error) {
	list := inst.getItems()
	for _, item := range list {
		if item.Mechanism.Support(a) {
			return item.Authorizer, nil
		}
	}
	return nil, fmt.Errorf("no Authorizer for Authorization, action:%s", a.Action())
}

func (inst *AuthorizerManagerImpl) getItems() []*auth.Registration {
	list := inst.cache
	if list == nil {
		list = inst.loadItems()
		inst.cache = list
	}
	return list
}

func (inst *AuthorizerManagerImpl) loadItems() []*auth.Registration {
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

func (inst *AuthorizerManagerImpl) accept(r *auth.Registration) bool {
	return r.Enabled && (r.Mechanism != nil) && (r.Authorizer != nil)
}
