package internal

import (
	"context"

	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

// AuthService2 ...
type AuthService2 struct {
	//starter:component
	_as func(rbac.AuthService) //starter:as("#")

	Servic1 auth.Service //starter:inject("#")
}

func (inst *AuthService2) _impl() rbac.AuthService {
	return inst
}

// Login 登录
func (inst *AuthService2) Login(c context.Context, a *rbac.AuthDTO) (*rbac.AuthDTO, error) {

	req := inst.prepare(a)

	err := inst.Servic1.Login(c, req)
	if err != nil {
		return nil, err
	}

	a2 := &rbac.AuthDTO{}
	a2.Mechanism = a.Mechanism
	a2.Account = a.Account
	return a2, nil
}

func (inst *AuthService2) prepare(src *rbac.AuthDTO) auth.Authentication {
	dst := &authService2request{}
	if src != nil {
		dst.data = *src
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type authService2request struct {
	data rbac.AuthDTO
}

func (inst *authService2request) Mechanism() string {
	return inst.data.Mechanism
}

func (inst *authService2request) Account() string {
	return inst.data.Account
}

func (inst *authService2request) Secret() []byte {
	return inst.data.Secret.Bytes()
}
