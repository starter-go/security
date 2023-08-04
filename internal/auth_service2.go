package internal

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
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

	req := inst.prepare(c, a)

	res, err := inst.Servic1.Login(c, req)
	if err != nil {
		return nil, err
	}

	a2 := &rbac.AuthDTO{}
	a2.Mechanism = a.Mechanism
	a2.Account = a.Account
	a2.Success = res.Success

	if res.Challenge {
		a2.Properties = make(map[string]string)
		a2.Properties["challenge"] = "yes"
	}

	return a2, nil
}

func (inst *AuthService2) prepare(ctx context.Context, src *rbac.AuthDTO) auth.Authentication {
	dst := &authService2request{}
	if src != nil {
		dst.data = *src
		dst.context = ctx
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type authService2request struct {
	data    rbac.AuthDTO
	context context.Context
	atts    attributes.Table
	params  parameters.Table
}

func (inst *authService2request) Context() context.Context {
	return inst.context
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

func (inst *authService2request) Attributes() attributes.Table {
	return inst.atts
}

func (inst *authService2request) Parameters() parameters.Table {
	return inst.params
}
