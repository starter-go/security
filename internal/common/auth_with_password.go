package common

import (
	"strings"

	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

// PasswordAuth ...
type PasswordAuth struct {

	//starter:component
	_as func(auth.Authenticator, auth.Registry) //starter:as(".",".")

	UserDao rbac.UserDAO //starter:inject("#")

}

func (inst *PasswordAuth) _impl() (auth.Authenticator, auth.Registry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *PasswordAuth) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Priority:      0,
		Enabled:       true,
		Authenticator: inst,
	}
	return []*auth.Registration{r1}
}

// Support ...
func (inst *PasswordAuth) Support(a auth.Authentication) bool {
	mech := a.Mechanism()
	mech = strings.TrimSpace(mech)
	mech = strings.ToLower(mech)
	return (mech == "password")
}

// Authenticate ...
func (inst *PasswordAuth) Authenticate(a auth.Authentication) (auth.User, error) {

	account := a.Account()
	user1, err := inst.UserDao.FindByAny(nil, account)
	if err != nil {
		return nil, err
	}

	target := user1.Password.Bytes()
	salt := user1.Salt.Bytes()
	pc := auth.PasswordCalculator{}
	pc.Init(target, salt)
	err = pc.Verify(a.Secret())
	if err != nil {
		return nil, err
	}

	ub := auth.UserBuilder{}
	ub.Avatar = user1.Avatar
	ub.DisplayName = user1.Nickname
	ub.ID = user1.ID
	ub.Name = user1.Name
	ub.Roles = user1.Roles

	user2 := ub.Create()
	return user2, nil
}
