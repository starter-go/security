package internal

import (
	"context"
	"time"

	"strings"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security/random"
)

// RbacSessionServiceImpl ...
type RbacSessionServiceImpl struct {

	//starter:component

	_as func(rbac.SessionService) //starter:as("#")

	KVS     keyvalues.Service  //starter:inject("#")
	UUIDGen random.UUIDService //starter:inject("#")

	MaxAge int64 //starter:inject("${security.session.default-age-sec}")

	class keyvalues.Class
}

func (inst *RbacSessionServiceImpl) _impl() rbac.SessionService { return inst }

func (inst *RbacSessionServiceImpl) getClass() (keyvalues.Class, error) {
	cl := inst.class
	if cl == nil {
		const (
			ns    = "github.com/starter-go/rbac"
			alias = "SessionDTO"
		)
		c2, err := inst.KVS.GetClassNS(ns, alias)
		if err != nil {
			return nil, err
		}
		cl = c2
		inst.class = cl
	}
	return cl, nil
}

func (inst *RbacSessionServiceImpl) model() *rbac.SessionDTO {
	return new(rbac.SessionDTO)
}

func (inst *RbacSessionServiceImpl) stringifySessionID(id rbac.SessionID) string {
	str := string(id)
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	if str == "" {
		str = "0000"
	}
	return str
}

func (inst *RbacSessionServiceImpl) getMaxAge() time.Duration {
	sec := inst.MaxAge
	ms := lang.Milliseconds(sec * 1000)
	return ms.Duration()
}

func (inst *RbacSessionServiceImpl) generateSessionID(se *rbac.SessionDTO) rbac.SessionID {
	ub := inst.UUIDGen.Build()
	ub = ub.Class("rbac.SessionDTO").ID(se.User.String())
	uuid := ub.Generate()
	return rbac.SessionID(uuid)
}

// Find ...
func (inst *RbacSessionServiceImpl) Find(c context.Context, id rbac.SessionID) (*rbac.SessionDTO, error) {

	mod := inst.model()
	idstr := inst.stringifySessionID(id)
	cl, err := inst.getClass()
	if err != nil {
		return nil, err
	}

	ent := cl.GetJSONEntry(idstr)
	err = ent.Get(mod)
	if err != nil {
		return nil, err
	}
	return mod, nil
}

// Insert ...
func (inst *RbacSessionServiceImpl) Insert(c context.Context, item *rbac.SessionDTO) (*rbac.SessionDTO, error) {

	id := inst.generateSessionID(item)
	idstr := inst.stringifySessionID(id)
	cl, err := inst.getClass()
	if err != nil {
		return nil, err
	}

	now := lang.Now()
	maxage := inst.getMaxAge()

	item.ID = id
	item.UUID = lang.UUID(id)
	item.CreatedAt = now
	item.UpdatedAt = now
	item.StartedAt = now
	item.ExpiredAt = now.Add(maxage)

	// put
	ent := cl.GetJSONEntry(idstr)
	opt := inst.makeOptions(item, now)
	err = ent.Put(item, opt)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// Update ...
func (inst *RbacSessionServiceImpl) Update(c context.Context, id rbac.SessionID, src *rbac.SessionDTO) (*rbac.SessionDTO, error) {

	dst := inst.model()
	idstr := inst.stringifySessionID(id)
	cl, err := inst.getClass()
	if err != nil {
		return nil, err
	}

	// fetch
	ent := cl.GetJSONEntry(idstr)
	err = ent.Get(dst)
	if err != nil {
		return nil, err
	}

	// update
	now := lang.Now()
	*dst = *src
	dst.UpdatedAt = now

	// vlog.Warn("todo: RbacSessionServiceImpl.Update() ... more fileds ")

	// save
	opt := inst.makeOptions(dst, now)
	err = ent.Put(dst, opt)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func (inst *RbacSessionServiceImpl) makeOptions(item *rbac.SessionDTO, now lang.Time) *keyvalues.Options {
	t1 := now
	t2 := item.ExpiredAt
	opt := new(keyvalues.Options)
	ttl := lang.Milliseconds(t2 - t1)
	if ttl > 0 {
		opt.MaxAge = ttl.Duration()
	} else {
		opt.MaxAge = inst.getMaxAge()
	}
	return opt
}
