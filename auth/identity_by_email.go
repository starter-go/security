package auth

import "github.com/starter-go/security/rbac"

// EmailIdentity 表示基于 email address 的身份信息
type EmailIdentity interface {
	Identity

	Address() rbac.EmailAddress
}

////////////////////////////////////////////////////////////////////////////////

// NewEmailIdentity ...
func NewEmailIdentity(by Authentication, info *rbac.EmailAddressDTO) EmailIdentity {
	id := &innerEmailIdentity{}
	id.info = *info
	id.mechanism = by.Mechanism()
	id.by = by
	return id
}

////////////////////////////////////////////////////////////////////////////////

type innerEmailIdentity struct {
	info      rbac.EmailAddressDTO
	mechanism string
	by        Authentication
}

func (inst *innerEmailIdentity) _impl() EmailIdentity {
	return inst
}

func (inst *innerEmailIdentity) By() Authentication {
	return inst.by
}

func (inst *innerEmailIdentity) Class() string {
	return "email"
}

func (inst *innerEmailIdentity) Name() string {
	return inst.info.Address.String()
}

func (inst *innerEmailIdentity) Mechanism() string {
	return inst.mechanism
}

func (inst *innerEmailIdentity) Address() rbac.EmailAddress {
	return inst.info.Address
}
