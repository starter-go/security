package permissions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// Dao ...
type Dao struct {
	//starter:component
	_as func(rbac.PermissionDAO) //starter:as("#")

	Agent         libgorm.Agent      //starter:inject("#")
	UUIDGenerator lang.UUIDGenerator //starter:inject("#")
}

func (inst *Dao) _impl() {
	inst._as(inst)
}

func (inst *Dao) model() *rbac.PermissionEntity {
	return &rbac.PermissionEntity{}
}

func (inst *Dao) modelList() []*rbac.PermissionEntity {
	return make([]*rbac.PermissionEntity, 0)
}

func (inst *Dao) makeResult(res *gorm.DB, o *rbac.PermissionEntity) (*rbac.PermissionEntity, error) {
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
	return inst.makeResult(res, o)
}

// List ...
func (inst *Dao) List(db *gorm.DB, q *rbac.PermissionQuery) ([]*rbac.PermissionEntity, error) {

	if q == nil {
		q = &rbac.PermissionQuery{}
		q.Pagination.Size = 10
	}

	offset := q.Pagination.Offset()
	limit := q.Pagination.Limit()
	list := inst.modelList()
	count := int64(0)

	db = inst.Agent.DB(db)
	db.Model(inst.model()).Count(&count)
	db = db.Offset(int(offset)).Limit(int(limit))
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}

	q.Pagination.Total = count
	return list, nil
}

// Insert ...
func (inst *Dao) Insert(db *gorm.DB, o *rbac.PermissionEntity) (*rbac.PermissionEntity, error) {

	uuid := inst.UUIDGenerator.Generate("rbac.PermissionEntity")

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(res, o)
}

// Update ...
func (inst *Dao) Update(db *gorm.DB, id rbac.PermissionID, updater func(*rbac.PermissionEntity)) (*rbac.PermissionEntity, error) {
	o := inst.model()
	db = inst.Agent.DB(db)
	res := db.Find(o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	updater(o)
	res = db.Save(o)
	return inst.makeResult(res, o)
}

// Delete ...
func (inst *Dao) Delete(db *gorm.DB, id rbac.PermissionID) error {
	db = inst.Agent.DB(db)
	res := db.Delete(inst.model(), id)
	return res.Error
}
