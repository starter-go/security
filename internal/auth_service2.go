package internal

import (
	"context"

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

// Handle 处理验证、授权请求
func (inst *AuthService2) Handle(c context.Context, a *rbac.AuthDTO) (*rbac.AuthDTO, error) {

	req := inst.prepare(c, a)

	err := inst.Servic1.Handle(req)
	if err != nil {
		return nil, err
	}

	params := req.Parameters().Export(nil)

	a2 := &rbac.AuthDTO{}
	a2.Mechanism = a.Mechanism
	a2.Account = a.Account
	a2.Parameters = params

	return a2, nil
}

func (inst *AuthService2) prepare(ctx context.Context, src *rbac.AuthDTO) auth.Authentication {

	ab := auth.RequestBuilder{
		Context:   ctx,
		Action:    src.Action,
		Mechanism: src.Mechanism,
		Account:   src.Account,
		Secret:    src.Secret,
	}

	params := parameters.NewTable(nil)
	params.Import(src.Parameters)
	ab.Params = params

	return ab.Create()
}

////////////////////////////////////////////////////////////////////////////////

// type authService2request struct {
// 	data    rbac.AuthDTO
// 	context context.Context
// 	atts    attributes.Table
// 	params  parameters.Table
// }

// func (inst *authService2request) _impl() auth.Authentication {
// 	return inst
// }

// func (inst *authService2request) Context() context.Context {
// 	return inst.context
// }

// func (inst *authService2request) Mechanism() string {
// 	return inst.data.Mechanism
// }

// func (inst *authService2request) Action() string {
// 	return inst.data.Action
// }

// func (inst *authService2request) Account() string {
// 	return inst.data.Account
// }

// func (inst *authService2request) Secret() []byte {
// 	return inst.data.Secret.Bytes()
// }

// func (inst *authService2request) Attributes() attributes.Table {
// 	return inst.atts
// }

// func (inst *authService2request) Parameters() parameters.Table {
// 	return inst.params
// }
