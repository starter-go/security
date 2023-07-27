package code

import (
	"context"
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/rbac"
	"github.com/starter-go/vlog"
)

// TestCom ...
type TestCom struct {

	//starter:component

	AuthSer       rbac.AuthService       //starter:inject("#")
	SessionSer    rbac.SessionService    //starter:inject("#")
	UserSer       rbac.UserService       //starter:inject("#")
	PermissionSer rbac.PermissionService //starter:inject("#")

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
	steps = append(steps, inst.doInsertUser)
	steps = append(steps, inst.doListUsers)
	steps = append(steps, inst.doInsertPermission)

	for _, fn := range steps {
		err := fn(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *TestCom) doLogin(c context.Context) error {

	password := "test"
	a1 := &rbac.AuthDTO{
		Mechanism: "password",
		Account:   "test",
		//	Secret:    password,
	}
	a1.Secret = lang.Base64FromBytes([]byte(password))

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

func (inst *TestCom) doInsertUser(c context.Context) error {

	now := lang.Now()
	userName := fmt.Sprintf("user_%d", now.Int())
	u2, err := inst.UserSer.Insert(c, &rbac.UserDTO{
		Name:     rbac.UserName(userName),
		NickName: "",
	})

	if err == nil {
		vlog.Debug("user.name = %s", u2.Name)
	}

	return err
}

func (inst *TestCom) doListUsers(c context.Context) error {
	list, err := inst.UserSer.List(c, nil)
	if err == nil {
		for i, item := range list {
			vlog.Debug("user[%d].name = %s", i, item.Name)
		}
	}
	return err
}

func (inst *TestCom) doInsertPermission(c context.Context) error {

	now := lang.Now()
	path := fmt.Sprintf("/a/b/c/%d", now.Int())
	u2, err := inst.PermissionSer.Insert(c, &rbac.PermissionDTO{
		Method: "PUT",
		Path:   path,
	})

	if err == nil {
		vlog.Debug("permission.path = %s", u2.Path)
	}

	return err
}
