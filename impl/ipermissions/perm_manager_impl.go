package ipermissions

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/security/permissions"
)

type PermissionManagerImpl struct {

	//starter:component

	_as func(permissions.Manager) //starter:as("#")

	Loader permissions.CacheLoader //starter:inject("#")

	cache permissions.Cache
}

// Find implements permissions.Manager.
func (inst *PermissionManagerImpl) Find(ctx context.Context, want *permissions.Perm) (*permissions.Perm, error) {
	c, err := inst.GetCache()
	if err != nil {
		return nil, err
	}
	return c.Find(ctx, want)
}

// GetCache implements permissions.Manager.
func (inst *PermissionManagerImpl) GetCache() (permissions.Cache, error) {

	c := inst.cache
	if c != nil {
		return c, nil
	}

	// load
	c, err := inst.innerLoadCache()
	if err != nil {
		return nil, err
	}
	inst.cache = c

	return c, nil
}

func (inst *PermissionManagerImpl) innerLoadCache() (permissions.Cache, error) {

	loader := inst.Loader
	c1, err := loader.Load()
	if err != nil {
		return nil, err
	}

	facade := new(innerCacheFacade)
	facade.man = inst
	facade.core = c1

	return facade, nil
}

// Life implements application.Lifecycle.
func (inst *PermissionManagerImpl) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.onCreate,
	}
}

func (inst *PermissionManagerImpl) onCreate() error {
	_, err := inst.GetCache()
	return err
}

func (inst *PermissionManagerImpl) _impl() (permissions.Manager, application.Lifecycle) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////

type innerCacheFacade struct {
	core permissions.Cache
	man  *PermissionManagerImpl
}

// Clear implements permissions.Cache.
func (inst *innerCacheFacade) Clear() {
	inst.man.cache = nil
}

// Find implements permissions.Cache.
func (inst *innerCacheFacade) Find(c context.Context, want *permissions.Perm) (*permissions.Perm, error) {
	return inst.core.Find(c, want)
}

func (inst *innerCacheFacade) _impl() permissions.Cache {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
