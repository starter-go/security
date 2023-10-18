package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/base/safe"
)

// Authorization 表示一个授权请求
type Authorization interface {
	Request

	Action() string

	Identities() []Identity
}

// Authorizer 表示一个授权组件
type Authorizer interface {

	// 向用户授权
	Authorize(a Authorization) error
}

// AuthorizerManager ...
type AuthorizerManager interface {
	FindFor(a Authorization) (Authorizer, error)
	ListFor(a Authorization) []Authorizer
}

////////////////////////////////////////////////////////////////////////////////

type authorization struct {
	context context.Context
	attrs   attributes.Table
	params  parameters.Table
	action  string
	ids     []Identity
}

func (inst *authorization) _impl() Authorization {
	return inst
}

func (inst *authorization) Context() context.Context {
	return inst.context
}

func (inst *authorization) Attributes() attributes.Table {
	return inst.attrs
}

func (inst *authorization) Parameters() parameters.Table {
	return inst.params
}

func (inst *authorization) Action() string {
	return inst.action
}

func (inst *authorization) Identities() []Identity {
	return inst.ids
}

////////////////////////////////////////////////////////////////////////////////

// AuthorizationBuilder  用来创建 Authorization
type AuthorizationBuilder struct {
	Context    context.Context
	Attributes attributes.Table
	Parameters parameters.Table
	Action     string
	Identities []Identity
}

// Create ...
func (inst *AuthorizationBuilder) Create() Authorization {

	ctx := inst.Context
	atts := inst.Attributes
	params := inst.Parameters
	action := inst.Action
	ids := inst.Identities

	if ctx == nil {
		ctx = context.Background()
	}

	if atts == nil {
		atts = attributes.NewTable(safe.Fast())
	}

	if params == nil {
		params = parameters.NewTable(safe.Fast())
	}

	if ids == nil {
		ids = make([]Identity, 0)
	}

	return &authorization{
		context: ctx,
		attrs:   atts,
		params:  params,
		action:  action,
		ids:     ids,
	}
}

////////////////////////////////////////////////////////////////////////////////
