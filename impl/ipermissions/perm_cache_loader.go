package ipermissions

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/permissions"
	"github.com/starter-go/vlog"
)

type PermissionCacheLoader struct {

	//starter:component

	_as func(permissions.CacheLoader) //starter:as("#")

	Source []permissions.Registry //starter:inject(".")

}

// Load implements permissions.CacheLoader.
func (inst *PermissionCacheLoader) Load() (permissions.Cache, error) {

	src := inst.Source
	loading := new(innerPermLoading)

	loading.init(inst)

	for _, r1 := range src {
		if r1 == nil {
			continue
		}
		tmp := r1.ListRegistrations()
		for _, r2 := range tmp {
			err := loading.add(r2)
			if err != nil {
				return nil, err
			}
		}
	}

	return loading.cache, nil
}

func (inst *PermissionCacheLoader) getRegularURI(p *permissions.Registration) permissions.URI {
	const (
		prefix = "uri:permission:"
	)
	method := p.Method
	path := p.Path
	u := prefix + path + "#" + method
	u = strings.ToLower(u)
	return permissions.URI(u)
}

func (inst *PermissionCacheLoader) _impl() permissions.CacheLoader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerPermLoading struct {
	loader *PermissionCacheLoader
	cache  *innerPermCacheCore
}

func (inst *innerPermLoading) init(l *PermissionCacheLoader) {
	c := new(innerPermCacheCore)
	c.init(l)
	inst.loader = l
	inst.cache = c
}

func (inst *innerPermLoading) add(item1 *permissions.Registration) error {

	if !inst.accept(item1) {
		return nil
	}

	item2 := new(permissions.Perm)
	item2.Registration = *item1
	inst.preparePerm(item2)

	return inst.cache.add(item2)
}

func (inst *innerPermLoading) accept(p *permissions.Registration) bool {

	if p == nil {
		return false
	}

	if p.Method == "" {
		return false
	}

	if p.Path == "" {
		return false
	}

	if !p.Enabled {
		return false
	}

	return true
}

func (inst *innerPermLoading) preparePerm(perm *permissions.Perm) {

	loader := inst.loader
	uri := loader.getRegularURI(&perm.Registration)

	roles := perm.Roles.List()
	rolesmap := make(map[rbac.RoleName]bool)
	for _, role := range roles {
		rolesmap[role] = true
	}

	perm.AcceptAdmin = rolesmap[rbac.RoleAdmin]
	perm.AcceptAnonym = rolesmap[rbac.RoleAnonym]
	perm.AcceptAny = rolesmap[rbac.RoleAny]
	perm.AcceptFriend = rolesmap[rbac.RoleFriend]
	perm.AcceptGuest = rolesmap[rbac.RoleGuest]
	perm.AcceptOwner = rolesmap[rbac.RoleOwner]
	perm.AcceptRoot = rolesmap[rbac.RoleRoot]
	perm.AcceptUser = rolesmap[rbac.RoleUser]

	perm.URI = uri
}

////////////////////////////////////////////////////////////////////////////////

type innerPermCacheCore struct {
	mu     sync.Mutex
	table  map[permissions.URI]*permissions.Perm
	loader *PermissionCacheLoader
}

// Clear implements permissions.Cache.
func (inst *innerPermCacheCore) Clear() {
	// nop
}

// Find implements permissions.Cache.
func (inst *innerPermCacheCore) Find(c context.Context, want *permissions.Perm) (*permissions.Perm, error) {

	if want == nil {
		return nil, fmt.Errorf("method param 'want' is nil")
	}

	mu := &inst.mu
	mu.Lock()
	defer mu.Unlock()

	table := inst.table
	loader := inst.loader
	uri := loader.getRegularURI(&want.Registration)
	have := table[uri]

	if have == nil {
		return nil, fmt.Errorf("no permissions.Perm with URI '%s'", uri)
	}

	*want = *have
	return want, nil
}

func (inst *innerPermCacheCore) add(it *permissions.Perm) error {

	uri := it.URI
	table := inst.table
	older := table[uri]
	sel := it

	if older != nil {
		if older.Priority < it.Priority {
			sel = it
		} else {
			sel = older
		}
		vlog.Warn("a permissions.Registration is duplicate, URI = '%s'", uri)
	}

	table[uri] = sel
	return nil
}

func (inst *innerPermCacheCore) init(l *PermissionCacheLoader) {
	inst.table = make(map[permissions.URI]*permissions.Perm)
	inst.loader = l
}

func (inst *innerPermCacheCore) _impl() permissions.Cache {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
