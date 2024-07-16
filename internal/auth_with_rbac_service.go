package internal

import (
	"context"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security/auth"
)

// AuthService2 是面向 rabc 的 auth 服务
type AuthService2 struct {
	//starter:component
	_as func(rbac.AuthService) //starter:as("#")

	Servic1 auth.Service //starter:inject("#")
}

func (inst *AuthService2) _impl() rbac.AuthService {
	return inst
}

// Handle 处理验证、授权请求
func (inst *AuthService2) Handle(c context.Context, action string, alist []*rbac.AuthDTO) ([]*rbac.AuthDTO, error) {
	fb := auth.NewFeedback()
	reqlist := inst.prepare(c, action, alist, fb)
	err := inst.Servic1.Execute(reqlist...)
	if err != nil {
		return nil, err
	}
	return inst.makeFeedbackResult(fb)
}

func (inst *AuthService2) makeFeedbackResult(fb auth.Feedback) ([]*rbac.AuthDTO, error) {

	src := fb.Parameters().Export(nil)
	item := &rbac.AuthDTO{}
	item.Parameters = src

	list := make([]*rbac.AuthDTO, 0)
	list = append(list, item)
	return list, nil
}

func (inst *AuthService2) prepare(ctx context.Context, action string, src []*rbac.AuthDTO, fb auth.Feedback) []auth.Request {

	mode := safe.Fast()
	c := &authService2context{}
	c.ctx = ctx
	c.att = attributes.NewTable(mode)
	result := make([]auth.Request, 0)
	step := ""

	for _, item := range src {
		c.params = parameters.NewTable(mode)
		inst.copyParams(item.Parameters, c.params)
		if item.Account != "" {
			// as authentication
			req := inst.prepareAuthentication(c, action, item, fb)
			result = append(result, req)
		}
		if item.Step != "" {
			step = item.Step
		}
	}

	// make authorization
	req := inst.prepareAuthorization(c, action, step, fb)
	result = append(result, req)

	return result
}

func (inst *AuthService2) copyParams(src map[string]string, dst parameters.Table) {
	for k, v := range src {
		dst.SetParam(k, v)
	}
}

func (inst *AuthService2) prepareAuthentication(c *authService2context, action string, src *rbac.AuthDTO, fb auth.Feedback) auth.Authentication {
	ab := auth.AuthenticationBuilder{
		Attributes: c.att,
		Context:    c.ctx,
		Parameters: c.params,
		Mechanism:  src.Mechanism,
		Account:    src.Account,
		Secret:     src.Secret,
		Action:     action,
		Step:       src.Step,
		Feedback:   fb,
	}
	return ab.Create()
}

func (inst *AuthService2) prepareAuthorization(c *authService2context, action string, step string, fb auth.Feedback) auth.Authorization {
	ab := auth.AuthorizationBuilder{
		Attributes: c.att,
		Context:    c.ctx,
		Parameters: c.params,
		Action:     action,
		Step:       step,
		Identities: nil,
		Feedback:   fb,
	}
	return ab.Create()
}

////////////////////////////////////////////////////////////////////////////////

type authService2context struct {
	ctx    context.Context
	att    attributes.Table
	params parameters.Table
}
