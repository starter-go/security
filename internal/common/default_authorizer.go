package common

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/jwt"
)

// DefaultAuthorizer 是默认的授权组件
type DefaultAuthorizer struct {

	//starter:component
	_as func(auth.Registry) //starter:as(".")

	Service jwt.Service //starter:inject("#")
	MaxAge  int         //starter:inject("${security.jwt.max-age}")
}

func (inst *DefaultAuthorizer) _impl() (auth.Authorizer, auth.Registry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DefaultAuthorizer) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Priority:   0,
		Enabled:    true,
		Authorizer: inst,
	}
	return []*auth.Registration{r1}
}

// Authorize ...
func (inst *DefaultAuthorizer) Authorize(a auth.Authorization) error {

	ctx := a.Context()
	user := a.User()
	now := lang.Now()

	o1 := &jwt.DTO{}
	user2 := &o1.Session.User

	user2.Avatar = user.Avatar()
	user2.ID = user.ID()
	user2.Name = user.Name()
	user2.NickName = user.DisplayName()
	user2.Roles = user.Roles()

	o1.CreatedAt = now
	o1.ExpiredAt = inst.computeExpiredAt(now)

	t1, err := inst.Service.Encode(o1)
	if err != nil {
		return err
	}
	return inst.Service.SetText(ctx, t1)
}

// Support ...
func (inst *DefaultAuthorizer) Support(a auth.Authorization) bool {
	return true
}

func (inst *DefaultAuthorizer) computeExpiredAt(now lang.Time) lang.Time {
	ma := lang.Time(inst.MaxAge)
	if ma < 0 {
		ma = 0
	}
	return now + ma
}
