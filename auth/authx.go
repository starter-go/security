package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/base/lang"
)

////////////////////////////////////////////////////////////////////////////////

// RequestBuilder 用于创建 Authentication & Authorization 实例
type RequestBuilder struct {
	Context   context.Context
	Atts      attributes.Table
	Params    parameters.Table
	Account   string
	Action    string
	Mechanism string
	Secret    lang.Base64
}

// Create 创建 Authentication 实例
func (inst *RequestBuilder) Create() Authentication {

	a := &innerAuthentication{
		ctx:       inst.Context,
		account:   inst.Account,
		action:    inst.Action,
		mechanism: inst.Mechanism,
		secret:    inst.Secret,
		atts:      inst.Atts,
		params:    inst.Params,
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

type innerAuthentication struct {
	ctx       context.Context
	account   string
	action    string
	mechanism string
	secret    lang.Base64
	atts      attributes.Table
	params    parameters.Table
}

func (inst *innerAuthentication) _impl() (Authentication, Authorization) {
	return inst, inst
}

func (inst *innerAuthentication) Context() context.Context {
	return inst.ctx
}

func (inst *innerAuthentication) Attributes() attributes.Table {
	return inst.atts
}

func (inst *innerAuthentication) Parameters() parameters.Table {
	return inst.params
}

func (inst *innerAuthentication) Action() string {
	return inst.action
}

func (inst *innerAuthentication) Mechanism() string {
	return inst.mechanism
}

func (inst *innerAuthentication) Account() string {
	return inst.account
}

func (inst *innerAuthentication) Secret() []byte {
	return inst.secret.Bytes()
}

////////////////////////////////////////////////////////////////////////////////
