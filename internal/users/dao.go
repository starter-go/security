package users

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// Dao ...
type Dao struct {
	//starter:component
	_as func(rbac.UserDAO) //starter:as("#")

	Agent libgorm.Agent //starter:inject("#")
}

func (inst *Dao) _impl() {
	inst._as(inst)
}

func (inst *Dao) model() *rbac.UserEntity {
	return &rbac.UserEntity{}
}

func (inst *Dao) makeReesult(res *gorm.DB, o *rbac.UserEntity) (*rbac.UserEntity, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *Dao) Find(db *gorm.DB, id rbac.UserID) (*rbac.UserEntity, error) {
	db = inst.Agent.DB(db)
	o := inst.model()
	res := db.Find(o, id)
	return inst.makeReesult(res, o)
}
