package internal

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/base/util"
	"github.com/starter-go/security/jwt"
)

// JWTService 实现 jwt.Service
type JWTService struct {
	//starter:component
	_as func(jwt.Service) //starter:as("#")

	RegistryList []jwt.Registry //starter:inject(".")

	TokenMaxageDefault int //starter:inject("${security.jwt.default-age-sec}")
	TokenMaxageMin     int //starter:inject("${security.jwt.min-age-sec}")
	TokenMaxageMax     int //starter:inject("${security.jwt.max-age-sec}")

	SessionMaxageDefault int //starter:inject("${security.session.default-age-sec}")
	SessionMaxageMin     int //starter:inject("${security.session.min-age-sec}")
	SessionMaxageMax     int //starter:inject("${security.session.max-age-sec}")

	// cache
	adapterList []jwt.Adapter
	codecList   []jwt.CODEC
	regList     []*jwt.Registration
}

func (inst *JWTService) _impl() jwt.Service {
	return inst
}

func (inst *JWTService) listRegs() []*jwt.Registration {
	dst := inst.regList
	if dst != nil {
		return dst
	}
	src := inst.RegistryList
	for _, item1 := range src {
		i2list := item1.ListRegistrations()
		for _, item2 := range i2list {
			if item2 == nil {
				continue
			}
			if !item2.Enabled {
				continue
			}
			dst = append(dst, item2)
		}
	}
	s := &util.Sorter{}
	s.OnLen = func() int { return len(dst) }
	s.OnLess = func(i1, i2 int) bool { return dst[i1].Priority < dst[i2].Priority }
	s.OnSwap = func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] }
	s.Sort()
	inst.regList = dst
	return dst
}

func (inst *JWTService) listCODECs() []jwt.CODEC {
	dst := inst.codecList
	if dst != nil {
		return dst
	}
	src := inst.listRegs()
	for _, item1 := range src {
		item2 := item1.CODEC
		if item2 == nil {
			continue
		}
		dst = append(dst, item2)
	}
	inst.codecList = dst
	return dst
}

func (inst *JWTService) listAdapters() []jwt.Adapter {
	dst := inst.adapterList
	if dst != nil {
		return dst
	}
	src := inst.listRegs()
	for _, item1 := range src {
		item2 := item1.Adapter
		if item2 == nil {
			continue
		}
		dst = append(dst, item2)
	}
	inst.adapterList = dst
	return dst
}

// SetDTO ...
func (inst *JWTService) SetDTO(c context.Context, o *jwt.Token) error {
	inst.prepareToken(o)
	err := fmt.Errorf("no jwt.adapter for the context")
	list := inst.listAdapters()
	for _, item := range list {
		if item.Accept(c) {
			err = item.SetDTO(c, o)
			if err == nil {
				return nil
			}
		}
	}
	return err
}

func (inst *JWTService) prepareToken(token *jwt.Token) {

	if token == nil {
		return
	}

	computer := &maxAgeComputer{}
	session := &token.Session

	// token
	computer.ageMax = lang.Seconds(inst.TokenMaxageMax)
	computer.ageMin = lang.Seconds(inst.TokenMaxageMin)
	computer.ageDefault = lang.Seconds(inst.TokenMaxageDefault)

	a, b, c := computer.compute(token.CreatedAt, token.MaxAge, token.ExpiredAt)
	token.CreatedAt = a
	token.MaxAge = b
	token.ExpiredAt = c

	// session
	computer.ageMax = lang.Seconds(inst.SessionMaxageMax)
	computer.ageMin = lang.Seconds(inst.SessionMaxageMin)
	computer.ageDefault = lang.Seconds(inst.SessionMaxageDefault)

	a, b, c = computer.compute(session.CreatedAt, session.MaxAge, session.ExpiredAt)
	session.CreatedAt = a
	session.MaxAge = b
	session.ExpiredAt = c
}

// SetText ...
func (inst *JWTService) SetText(c context.Context, t jwt.Text) error {
	err := fmt.Errorf("no jwt.adapter for the context")
	list := inst.listAdapters()
	for _, item := range list {
		if item.Accept(c) {
			err = item.SetText(c, t)
			if err == nil {
				return nil
			}
		}
	}
	return err
}

// GetDTO ...
func (inst *JWTService) GetDTO(c context.Context) (*jwt.Token, error) {
	err := fmt.Errorf("no jwt.adapter for the context")
	list := inst.listAdapters()
	for _, item := range list {
		if item.Accept(c) {
			o, err2 := item.GetDTO(c)
			if err2 == nil {
				return o, nil
			}
			err = err2
		}
	}
	return nil, err
}

// GetText ...
func (inst *JWTService) GetText(c context.Context) (jwt.Text, error) {
	err := fmt.Errorf("no jwt.adapter for the context")
	list := inst.listAdapters()
	for _, item := range list {
		if item.Accept(c) {
			t, err2 := item.GetText(c)
			if err2 == nil {
				return t, nil
			}
			err = err2
		}
	}
	return "", err
}

// Encode ...
func (inst *JWTService) Encode(o *jwt.Token) (jwt.Text, error) {
	err := fmt.Errorf("no jwt.CODEC")
	inst.prepareToken(o)
	list := inst.listCODECs()
	for _, item := range list {
		t, err2 := item.Encode(o)
		if err2 == nil {
			return t, nil
		}
		err = err2
	}
	return "", err
}

// Decode ...
func (inst *JWTService) Decode(t jwt.Text) (*jwt.Token, error) {
	err := fmt.Errorf("no jwt.CODEC")
	list := inst.listCODECs()
	for _, item := range list {
		o, err2 := item.Decode(t)
		if err2 == nil {
			return o, nil
		}
		err = err2
	}
	return nil, err
}
