package testcom

import (
	"net/http"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/permissions"
)

type MockPermRegDemo struct {

	//starter:component

	_as func(permissions.Registry) //starter:as(".")

}

// ListRegistrations implements permissions.Registry.
func (inst *MockPermRegDemo) ListRegistrations() []*permissions.Registration {

	r1 := new(permissions.Registration)
	r1.Priority = 1
	r1.Method = http.MethodGet
	r1.Path = "/mock/demo1"
	r1.Roles = rbac.NewRoleNameList(rbac.RoleUser)
	r1.Enabled = true

	r2 := new(permissions.Registration)
	r2.Priority = 2
	r2.Method = http.MethodPost
	r2.Path = "/mock/demo1"
	r2.Roles = rbac.NewRoleNameList(rbac.RoleAdmin, rbac.RoleOwner, rbac.RoleRoot)
	r2.Enabled = true

	list := make([]*permissions.Registration, 0)
	list = append(list, r1)
	list = append(list, r2)
	return list
}

func (inst *MockPermRegDemo) _impl() permissions.Registry {
	return inst
}
