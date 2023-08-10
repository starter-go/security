package internal

import (
	"fmt"

	"github.com/starter-go/base/util"
	"github.com/starter-go/security/auth"
)

// AuthService1 ...
type AuthService1 struct {
	//starter:component
	_as func(auth.Service) //starter:as("#")

	RegistryList []auth.Registry //starter:inject(".")

	组件注册信息list []*auth.Registration
}

func (inst *AuthService1) _impl() auth.Service {
	return inst
}

// Handle 登录
func (inst *AuthService1) Handle(r auth.Authentication) error {

	reg, err := inst.findMechanism(r)
	if err != nil {
		return err
	}
	err = reg.Authenticator.Authenticate(r)
	if err != nil {
		return err
	}
	return reg.Authorizer.Authorize(r)
}

func (inst *AuthService1) getRegistrationList() []*auth.Registration {
	list := inst.组件注册信息list
	if list == nil {
		list = inst.loadRegistrationList()
		inst.组件注册信息list = list
	}
	return list
}

func (inst *AuthService1) loadRegistrationList() []*auth.Registration {

	dst := make([]*auth.Registration, 0)
	src := inst.RegistryList

	for _, r1 := range src {
		r2list := r1.ListRegistrations()
		for _, r2 := range r2list {
			if inst.isRegistrationOK(r2) {
				dst = append(dst, r2)
			}
		}
	}

	list := dst
	s := &util.Sorter{
		OnLen:  func() int { return len(list) },
		OnLess: func(i1, i2 int) bool { return list[i1].Priority < list[i2].Priority },
		OnSwap: func(i1, i2 int) { list[i1], list[i2] = list[i2], list[i1] },
	}
	s.Sort()
	return dst
}

func (inst *AuthService1) isRegistrationOK(r *auth.Registration) bool {

	if r == nil {
		return false
	}

	if !r.Enabled {
		return false
	}

	if r.Mechanism == nil || r.Authenticator == nil || r.Authorizer == nil {
		return false
	}

	return true
}

func (inst *AuthService1) findMechanism(r auth.Request) (*auth.Registration, error) {
	src := inst.getRegistrationList()
	for _, reg := range src {
		if reg.Mechanism.Support(r) {
			return reg, nil
		}
	}
	action := r.Action()
	mech := r.Mechanism()
	const f = "no mechanism for the authx request (Mechanism:%s, Action:%s)"
	return nil, fmt.Errorf(f, mech, action)
}
