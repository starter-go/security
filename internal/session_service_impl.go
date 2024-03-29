package internal

import (
	"context"
	"fmt"

	"github.com/starter-go/base/util"
	"github.com/starter-go/security"
)

// SessionServiceImpl ...
type SessionServiceImpl struct {

	//starter:component
	_as func(security.SessionService) //starter:as("#")

	RegistryList []security.SessionRegistry //starter:inject(".")

	providers []*security.SessionRegistration
}

func (inst *SessionServiceImpl) _impl() security.SessionService {
	return inst
}

// GetCurrent ...
func (inst *SessionServiceImpl) GetCurrent(c context.Context) (security.Session, error) {
	all := inst.getProviders()
	for _, r := range all {
		p := r.Provider
		if p.Support(c) {
			return p.Current(c)
		}
	}
	return nil, fmt.Errorf("no session provider for the context")
}

func (inst *SessionServiceImpl) getProviders() []*security.SessionRegistration {
	list := inst.providers
	if list == nil {
		list = inst.loadProviders()
		inst.providers = list
	}
	return list
}

func (inst *SessionServiceImpl) loadProviders() []*security.SessionRegistration {

	src := inst.RegistryList
	dst := make([]*security.SessionRegistration, 0)

	for _, r1 := range src {
		r2 := r1.Registration()
		if inst.isRegistrationOK(r2) {
			dst = append(dst, r2)
		}
	}

	s := &util.Sorter{}
	s.OnLen = func() int { return len(dst) }
	s.OnLess = func(i1, i2 int) bool { return dst[i1].Priority < dst[i2].Priority }
	s.OnSwap = func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] }
	s.Sort()

	return dst
}

func (inst *SessionServiceImpl) isRegistrationOK(r *security.SessionRegistration) bool {

	if r == nil {
		return false
	}

	if !r.Enabled {
		return false
	}

	if r.Provider == nil {
		return false
	}

	return true
}
