package auth

import "github.com/starter-go/rbac"

// PhoneIdentity 表示基于 phone number 的身份信息
type PhoneIdentity interface {
	Identity

	FullNumber() rbac.FullPhoneNumber
}

////////////////////////////////////////////////////////////////////////////////

// NewPhoneIdentity ...
func NewPhoneIdentity(by Authentication, info *rbac.PhoneNumberDTO) PhoneIdentity {
	id := &innerPhoneIdentity{}
	id.info = *info
	id.mechanism = by.Mechanism()
	id.by = by
	return id
}

////////////////////////////////////////////////////////////////////////////////

type innerPhoneIdentity struct {
	mechanism string
	info      rbac.PhoneNumberDTO
	by        Authentication
}

func (inst *innerPhoneIdentity) _impl() PhoneIdentity {
	return inst
}

func (inst *innerPhoneIdentity) Mechanism() string {
	return inst.mechanism
}

func (inst *innerPhoneIdentity) By() Authentication {
	return inst.by
}

func (inst *innerPhoneIdentity) Class() string {
	return "sms"
}

func (inst *innerPhoneIdentity) Name() string {
	return inst.info.FullNumber.String()
}

func (inst *innerPhoneIdentity) FullNumber() rbac.FullPhoneNumber {
	return inst.info.FullNumber
}
