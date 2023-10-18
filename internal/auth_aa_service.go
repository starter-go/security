package internal

import (
	"github.com/starter-go/security/auth"
)

// AuthService1 ... 是面向 a&a 的 auth 服务
type AuthService1 struct {
	//starter:component
	_as func(auth.Service) //starter:as("#")

	Authenticators auth.AuthenticatorManager //starter:inject("#")
	Authorizers    auth.AuthorizerManager    //starter:inject("#")
}

func (inst *AuthService1) _impl() auth.Service {
	return inst
}

// Authenticate 验证
func (inst *AuthService1) Authenticate(a auth.Authentication) ([]auth.Identity, error) {
	handler, err := inst.Authenticators.FindFor(a)
	if err != nil {
		return nil, err
	}
	return handler.Authenticate(a)
}

// Authorize 授权
func (inst *AuthService1) Authorize(a auth.Authorization) error {
	handler, err := inst.Authorizers.FindFor(a)
	if err != nil {
		return err
	}
	return handler.Authorize(a)
}

// Execute 登录
func (inst *AuthService1) Execute(reqlist ...auth.Request) error {
	alist1 := make([]auth.Authentication, 0)
	alist2 := make([]auth.Authorization, 0)
	for _, req := range reqlist {
		a1, ok := req.(auth.Authentication)
		if ok {
			alist1 = append(alist1, a1)
		}
		a2, ok := req.(auth.Authorization)
		if ok {
			alist2 = append(alist2, a2)
		}
	}
	ids, err := inst.authenticateAll(alist1)
	if err != nil {
		return err
	}
	return inst.authorizeAll(alist2, ids)
}

// 批量验证
func (inst *AuthService1) authenticateAll(todolist []auth.Authentication) ([]auth.Identity, error) {
	result := make([]auth.Identity, 0)
	for _, a1 := range todolist {
		ids, err := inst.Authenticate(a1)
		if err != nil {
			return nil, err
		}
		result = append(result, ids...)
	}
	return result, nil
}

// 批量授权
func (inst *AuthService1) authorizeAll(todolist []auth.Authorization, ids []auth.Identity) error {
	for _, a21 := range todolist {
		a22 := inst.makeAuthorizationWithIDs(ids, a21)
		err := inst.Authorize(a22)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *AuthService1) makeAuthorizationWithIDs(ids []auth.Identity, raw auth.Authorization) auth.Authorization {
	ab := auth.AuthorizationBuilder{}
	ab.Action = raw.Action()
	ab.Attributes = raw.Attributes()
	ab.Parameters = raw.Parameters()
	ab.Context = raw.Context()
	ab.Identities = ids
	return ab.Create()
}
