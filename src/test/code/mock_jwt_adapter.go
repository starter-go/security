package code

import (
	"context"

	"github.com/starter-go/security/jwt"
	"github.com/starter-go/vlog"
)

// MockTokenAdapter ...
type MockTokenAdapter struct {
	//starter:component
	_as func(jwt.Adapter, jwt.Registry) //starter:as(".",".")

	Service jwt.Service //starter:inject("#")
	Logger  vlog.Logger //starter:inject("#")
}

func (inst *MockTokenAdapter) _impl() (jwt.Adapter, jwt.Registry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *MockTokenAdapter) ListRegistrations() []*jwt.Registration {
	r1 := &jwt.Registration{
		Enabled:  true,
		Priority: 0,
		Adapter:  inst,
	}
	return []*jwt.Registration{r1}
}

// Accept ...
func (inst *MockTokenAdapter) Accept(c context.Context) bool {
	return true
}

// GetDTO ...
func (inst *MockTokenAdapter) GetDTO(c context.Context) (*jwt.Token, error) {
	t1, err := inst.GetText(c)
	if err != nil {
		return nil, err
	}
	return inst.Service.Decode(t1)
}

// GetText ...
func (inst *MockTokenAdapter) GetText(c context.Context) (jwt.Text, error) {
	return "23333333333333333333333333", nil
}

// SetDTO ...
func (inst *MockTokenAdapter) SetDTO(c context.Context, o1 *jwt.Token) error {
	t2, err := inst.Service.Encode(o1)
	if err != nil {
		return err
	}
	return inst.SetText(c, t2)
}

// SetText ...
func (inst *MockTokenAdapter) SetText(c context.Context, t1 jwt.Text) error {
	inst.Logger.Info("set-jwt: %s", t1.String())
	return nil
}
