package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/base/lang"
)

// Authentication 表示一个身份验证请求
type Authentication interface {
	Context() context.Context
	Attributes() attributes.Table
	Parameters() parameters.Table
	Action() string
	Mechanism() string
	Account() string
	Secret() []byte
}

// Authenticator 表示一个身份验证算法
type Authenticator interface {

	// 验证用户身份
	Authenticate(a Authentication) (*Result, error)

	Support(a Authentication) bool
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

func (inst *innerAuthentication) _impl() Authentication {
	return inst
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

// AuthenticationBuilder 用于创建 Authentication 实例
type AuthenticationBuilder struct {
	ctx       context.Context
	account   string
	action    string
	mechanism string
	secret    lang.Base64
	atts      attributes.Table
	params    parameters.Table
}

// Create 创建 Authentication 实例
func (inst *AuthenticationBuilder) Create() Authentication {

	a := &innerAuthentication{
		ctx:       inst.ctx,
		account:   inst.account,
		action:    inst.action,
		mechanism: inst.mechanism,
		secret:    inst.secret,
		atts:      inst.atts,
		params:    inst.params,
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
