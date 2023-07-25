package code

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/security/rbac"
)

// TestCom ...
type TestCom struct {

	//starter:component

	AuthSer    rbac.AuthService    //starter:inject("#")
	SessionSer rbac.SessionService //starter:inject("#")

}

// Life ...
func (inst *TestCom) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.test,
	}
}

func (inst *TestCom) test() error {

	c := context.Background()
	steps := make([]func(context.Context) error, 0)

	steps = append(steps, inst.doLogin)
	steps = append(steps, inst.doCurrentSession)

	for _, fn := range steps {
		err := fn(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *TestCom) doLogin(c context.Context) error {

	a1 := &rbac.AuthDTO{}

	a2, err := inst.AuthSer.Login(c, a1)
	if err != nil {
		return err
	}

	a2.DeletedAt.Int() // todo ...
	return nil
}

func (inst *TestCom) doCurrentSession(c context.Context) error {

	ses, err := inst.SessionSer.GetCurrent(c)
	if err != nil {
		return err
	}

	ses.DeletedAt.Int() // todo ...
	return nil
}
