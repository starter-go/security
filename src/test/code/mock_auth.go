package code

import "github.com/starter-go/security/auth"

// MockAuth 是虚拟的身份验证器
type MockAuth struct {

	//starter:component
	_as func(auth.Authenticator, auth.Registry) //starter:as(".",".")

}

func (inst *MockAuth) _impl() (auth.Authenticator, auth.Registry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *MockAuth) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Priority:      0,
		Enabled:       true,
		Authenticator: inst,
	}
	return []*auth.Registration{r1}
}

// Authenticate ...
func (inst *MockAuth) Authenticate(a auth.Authentication) (*auth.Result, error) {
	ub := &auth.UserBuilder{
		ID:          10000,
		Name:        "mock",
		DisplayName: "Mock",
		Roles:       "mock,user,admin,test",
		Avatar:      "https://example.com/mock/mock.png",
	}
	user := ub.Create()
	res := &auth.Result{
		User:      user,
		Success:   false,
		Challenge: true,
	}
	return res, nil
}

// Support ...
func (inst *MockAuth) Support(a auth.Authentication) bool {
	return a.Mechanism() == "mock"
}
