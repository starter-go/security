package auth

import "github.com/starter-go/security/rbac"

// PhoneIdentity 表示基于 phone number 的身份信息
type PhoneIdentity interface {
	Identity

	FullNumber() rbac.FullPhoneNumber
}

////////////////////////////////////////////////////////////////////////////////

// NewPhoneIdentity ...
func NewPhoneIdentity(mechanism string, info *rbac.PhoneNumberDTO) PhoneIdentity {
	id := &innerPhoneIdentity{}
	id.info = *info
	id.mechanism = mechanism
	return id
}

////////////////////////////////////////////////////////////////////////////////

type innerPhoneIdentity struct {
	mechanism string
	info      rbac.PhoneNumberDTO
}

func (inst *innerPhoneIdentity) _impl() PhoneIdentity {
	return inst
}

func (inst *innerPhoneIdentity) Mechanism() string {
	return inst.mechanism
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
