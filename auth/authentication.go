package auth

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/base/safe"
)

// Authentication 表示一个身份验证请求
type Authentication interface {
	Request

	Mechanism() string
	Account() string
	Secret() []byte
}

// Authenticator 表示一个身份验证算法
type Authenticator interface {

	// 验证用户身份
	Authenticate(a Authentication) ([]Identity, error)
}

// AuthenticatorManager ...
type AuthenticatorManager interface {
	FindFor(a Authentication) (Authenticator, error)
	ListFor(a Authentication) []Authenticator
}

////////////////////////////////////////////////////////////////////////////////

// AuthenticationBuilder 用于创建 Authentication 实例
type AuthenticationBuilder struct {
	Context    context.Context
	Attributes attributes.Table
	Parameters parameters.Table
	Account    string
	Mechanism  string
	Secret     lang.Base64
}

// Create 创建 Authentication 实例
func (inst *AuthenticationBuilder) Create() Authentication {

	a := &innerAuthentication{
		ctx:       inst.Context,
		account:   inst.Account,
		mechanism: inst.Mechanism,
		secret:    inst.Secret,
		atts:      inst.Attributes,
		params:    inst.Parameters,
	}

	if a.ctx == nil {
		a.ctx = context.Background()
	}

	if a.atts == nil {
		a.atts = attributes.NewTable(safe.Fast())
	}

	if a.params == nil {
		a.params = parameters.NewTable(safe.Fast())
	}

	return a
}

////////////////////////////////////////////////////////////////////////////////

type innerAuthentication struct {
	ctx       context.Context
	account   string
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
