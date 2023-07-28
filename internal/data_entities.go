package internal

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/security/rbac"
)

// TableReg ...
type TableReg struct {
	//starter:component
	_as func(libgorm.TableRegistry) //starter:as(".")
}

func (inst *TableReg) _impl() { inst._as(inst) }

// ListTableRegistrations ...
func (inst *TableReg) ListTableRegistrations() []*libgorm.TableRegistration {

	group := rbac.GetTableGroup()
	builder := &libgorm.TableGroupBuilder{}

	builder.Group(&libgorm.TableGroup{
		Name:        group.Name(),
		Namespace:   group.Namespace(),
		NamerSetter: group.SetNamer,
	})

	list := group.Entities()
	for _, ent := range list {
		builder.Entity(ent)
	}

	return builder.Create()
}
