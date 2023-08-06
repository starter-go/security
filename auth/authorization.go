package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
)

// Authorization 表示一个授权请求
type Authorization interface {
	Context() context.Context
	Attributes() attributes.Table
	Parameters() parameters.Table
	User() User
	Action() string
	Mechanism() string
}

// Authorizer 表示一个授权组件
type Authorizer interface {
	Authorize(a Authorization) error
	Support(a Authorization) bool
}

////////////////////////////////////////////////////////////////////////////////

type innerAuthorization struct {
	ctx       context.Context
	user      User
	action    string
	mechanism string
	atts      attributes.Table
	params    parameters.Table
}

func (inst *innerAuthorization) Context() context.Context {
	return inst.ctx
}

func (inst *innerAuthorization) User() User {
	return inst.user
}

func (inst *innerAuthorization) Action() string {
	return inst.action
}

func (inst *innerAuthorization) Mechanism() string {
	return inst.mechanism
}

func (inst *innerAuthorization) Attributes() attributes.Table {
	return inst.atts
}

func (inst *innerAuthorization) Parameters() parameters.Table {
	return inst.params
}

////////////////////////////////////////////////////////////////////////////////

// NewAuthorization 新建一个 Authorization 的实例
// func NewAuthorization(c context.Context, user User) Authorization {
// 	return &innerAuthorization{
// 		ctx:  c,
// 		user: user,
// 	}
// }

// AuthorizationBuilder ... 用于创建 Authorization
type AuthorizationBuilder struct {
	Context    context.Context
	Attributes attributes.Table
	Parameters parameters.Table
	User       User
	Action     string
}

// Create ...
func (inst *AuthorizationBuilder) Create() Authorization {
	a := &innerAuthorization{
		ctx:    inst.Context,
		user:   inst.User,
		atts:   inst.Attributes,
		action: inst.Action,
	}

	if a.ctx == nil {
		a.ctx = context.Background()
	}

	if a.atts == nil {
		a.atts = attributes.NewTable(nil)
	}

	if a.params == nil {
		a.params = parameters.NewTable(nil)
	}

	return a
}

////////////////////////////////////////////////////////////////////////////////
