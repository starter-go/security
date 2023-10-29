package rbac

import (
	"context"
)

// PhoneNumberID ...
type PhoneNumberID int64

// PhoneNumber 表示电话号码
type PhoneNumber string

// SimplePhoneNumber 表示简短的电话号码
type SimplePhoneNumber PhoneNumber

// FullPhoneNumber 表示完整的电话号码
type FullPhoneNumber PhoneNumber

// PhoneNumberDTO ...
type PhoneNumberDTO struct {
	ID PhoneNumberID `json:"id"`

	BaseDTO

	RegionCode2  RegionPhoneCode   `json:"region"`
	SimpleNumber SimplePhoneNumber `json:"simple_number"`
	FullNumber   FullPhoneNumber   `json:"full_number"`
}

// PhoneNumberQuery 查询参数
type PhoneNumberQuery struct {
	Pagination Pagination
	Conditions map[string]string
	All        bool // 查询全部条目
}

// PhoneNumberService ...
type PhoneNumberService interface {
	Insert(c context.Context, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Update(c context.Context, id PhoneNumberID, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Delete(c context.Context, id PhoneNumberID) error

	Find(c context.Context, id PhoneNumberID) (*PhoneNumberDTO, error)
	List(c context.Context, q *PhoneNumberQuery) ([]*PhoneNumberDTO, error)
}

////////////////////////////////////////////////////////////////////////////////

func (num PhoneNumber) String() string {
	return string(num)
}

func (num SimplePhoneNumber) String() string {
	return string(num)
}

func (num FullPhoneNumber) String() string {
	return string(num)
}
