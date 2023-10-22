package auth

import "github.com/starter-go/security/rbac"

// UserIdentity ...
type UserIdentity interface {
	Identity

	UserName() rbac.UserName

	UserID() rbac.UserID

	Nickname() string

	Avatar() string

	Roles() rbac.RoleNameList
}

////////////////////////////////////////////////////////////////////////////////

// NewUserIdentity ...
func NewUserIdentity(mechanism string, info *rbac.UserDTO) UserIdentity {
	id := &innerUserIdentity{}
	id.userinfo = *info
	id.mechanism = mechanism
	return id
}

////////////////////////////////////////////////////////////////////////////////

// innerUserIdentity ...
type innerUserIdentity struct {
	mechanism string
	userinfo  rbac.UserDTO
}

func (inst *innerUserIdentity) _impl() UserIdentity {
	return inst
}

// Class ...
func (inst *innerUserIdentity) Class() string {
	return "appuser"
}

// Name ...
func (inst *innerUserIdentity) Name() string {
	return inst.userinfo.Name.String()
}

func (inst *innerUserIdentity) Mechanism() string {
	return inst.mechanism
}

// Roles ...
func (inst *innerUserIdentity) Roles() rbac.RoleNameList {
	return inst.userinfo.Roles
}

// UserID ...
func (inst *innerUserIdentity) UserID() rbac.UserID {
	return inst.userinfo.ID
}

// UserName ...
func (inst *innerUserIdentity) UserName() rbac.UserName {
	return inst.userinfo.Name
}

// Avatar ...
func (inst *innerUserIdentity) Avatar() string {
	return inst.userinfo.Avatar
}

// Nickname ...
func (inst *innerUserIdentity) Nickname() string {
	return inst.userinfo.NickName
}