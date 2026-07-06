package permissions

import (
	"context"
)

type Cache interface {
	Clear()
	Find(c context.Context, want *Perm) (*Perm, error)
}

type CacheLoader interface {
	Load() (Cache, error)
}

type Manager interface {
	GetCache() (Cache, error)
	Find(c context.Context, want *Perm) (*Perm, error)
}
