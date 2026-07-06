package permissions

import (
	"github.com/starter-go/rbac"
)

type DTO = rbac.PermissionDTO

type Perm struct {
	Registration

	AcceptAdmin  bool
	AcceptAnonym bool
	AcceptAny    bool
	AcceptFriend bool
	AcceptGuest  bool
	AcceptOwner  bool
	AcceptRoot   bool
	AcceptUser   bool

	URI URI // method+path
}
