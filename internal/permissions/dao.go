package permissions

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// Dao ...
type Dao struct {
	//starter:component
	_as func(rbac.PermissionDAO) //starter:as("#")

	Agent libgorm.Agent //starter:inject("#")
}

func (inst *Dao) _impl() {
	inst._as(inst)
}

func (inst *Dao) model() *rbac.PermissionEntity {
	return &rbac.PermissionEntity{}
}

func (inst *Dao) makeReesult(res *gorm.DB, o *rbac.PermissionEntity) (*rbac.PermissionEntity, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *Dao) Find(db *gorm.DB, id rbac.PermissionID) (*rbac.PermissionEntity, error) {
	db = inst.Agent.DB(db)
	o := inst.model()
	res := db.Find(o, id)
	return inst.makeReesult(res, o)
}
