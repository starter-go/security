package sessions

// // Dao ...
// type Dao struct {
// 	//starter:component
// 	_as func(rbac.RoleDAO) //starter:as("#")

// 	Agent libgorm.Agent //starter:inject("#")
// }

// func (inst *Dao) _impl()  rbac.Session {
// 	return  inst
// }

// func (inst *Dao) model() *rbac.RoleEntity {
// 	return &rbac.RoleEntity{}
// }

// func (inst *Dao) makeReesult(res *gorm.DB, o *rbac.RoleEntity) (*rbac.RoleEntity, error) {
// 	err := res.Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return o, nil
// }

// // Find ...
// func (inst *Dao) Find(db *gorm.DB, id rbac.RoleID) (*rbac.RoleEntity, error) {
// 	db = inst.Agent.DB(db)
// 	o := inst.model()
// 	res := db.Find(o, id)
// 	return inst.makeReesult(res, o)
// }
