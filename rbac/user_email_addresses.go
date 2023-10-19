package rbac

import (
	"context"
)

// EmailAddressID ...
type EmailAddressID int64

// EmailAddress 表示 'user@domain' 形式的邮件地址
type EmailAddress string

// EmailAddressDTO ...
type EmailAddressDTO struct {
	ID EmailAddressID `json:"id"`

	BaseDTO

	Address EmailAddress `json:"address"`
}

// EmailAddressQuery 查询参数
type EmailAddressQuery struct {
	Pagination Pagination
}

// EmailAddressService ...
type EmailAddressService interface {
	Insert(c context.Context, o *EmailAddressDTO) (*EmailAddressDTO, error)
	Update(c context.Context, id EmailAddressID, o *EmailAddressDTO) (*EmailAddressDTO, error)
	Delete(c context.Context, id EmailAddressID) error

	Find(c context.Context, id EmailAddressID) (*EmailAddressDTO, error)
	List(c context.Context, q *EmailAddressQuery) ([]*EmailAddressDTO, error)
}

////////////////////////////////////////////////////////////////////////////////

func (addr EmailAddress) String() string {
	return string(addr)
}
