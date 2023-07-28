package internal

import (
	"context"
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
	验证组件list   []auth.Authenticator
	授权组件list   []auth.Authorizer
}

func (inst *AuthService1) _impl() auth.Service {
	return inst
}

// Login 登录
func (inst *AuthService1) Login(c context.Context, a1 auth.Authentication) error {
	user, err := inst.Authenticate(a1)
	if err != nil {
		return err
	}
	a2 := auth.NewAuthorization(c, user)
	return inst.Authorize(a2)
}

func (inst *AuthService1) getRegistrationList() []*auth.Registration {

	list := inst.组件注册信息list
	if list != nil {
		return list
	}

	src := inst.RegistryList
	for _, r1 := range src {
		r2list := r1.ListRegistrations()
		for _, r2 := range r2list {
			if r2 == nil {
				continue
			}
			if !r2.Enabled {
				continue
			}
			list = append(list, r2)
		}
	}

	s := &util.Sorter{
		OnLen:  func() int { return len(list) },
		OnLess: func(i1, i2 int) bool { return list[i1].Priority < list[i2].Priority },
		OnSwap: func(i1, i2 int) { list[i1], list[i2] = list[i2], list[i1] },
	}
	s.Sort()

	inst.组件注册信息list = list
	return list
}

func (inst *AuthService1) getAuthenticatorList() []auth.Authenticator {
	list := inst.验证组件list
	if list != nil {
		return list
	}
	src := inst.getRegistrationList()
	for _, reg := range src {
		if reg.Authenticator != nil {
			list = append(list, reg.Authenticator)
		}
	}
	inst.验证组件list = list
	return list
}

func (inst *AuthService1) getAuthorizerList() []auth.Authorizer {
	list := inst.授权组件list
	if list != nil {
		return list
	}
	src := inst.getRegistrationList()
	for _, reg := range src {
		if reg.Authorizer != nil {
			list = append(list, reg.Authorizer)
		}
	}
	inst.授权组件list = list
	return list
}

// Authorize 授权
func (inst *AuthService1) Authorize(a auth.Authorization) error {
	list := inst.getAuthorizerList()
	for _, item := range list {
		if item.Support(a) {
			return item.Authorize(a)
		}
	}
	return fmt.Errorf("bad auth")
}

// Authenticate 验证
func (inst *AuthService1) Authenticate(a auth.Authentication) (auth.User, error) {
	list := inst.getAuthenticatorList()
	for _, item := range list {
		if item.Support(a) {
			user, err := item.Authenticate(a)
			if err == nil {
				return user, nil
			}
		}
	}
	return nil, fmt.Errorf("bad auth")
}
