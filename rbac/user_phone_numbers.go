package rbac

import (
	"context"
)

// PhoneNumberID ...
type PhoneNumberID int64

// PhoneNumberDTO ...
type PhoneNumberDTO struct {
	ID PhoneNumberID `json:"id"`

	BaseDTO

	SimpleNunber string `json:"simple_number"`
}

// PhoneNumberQuery 查询参数
type PhoneNumberQuery struct {
	Pagination Pagination
}

// PhoneNumberService ...
type PhoneNumberService interface {
	Insert(c context.Context, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Update(c context.Context, id PhoneNumberID, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Delete(c context.Context, id PhoneNumberID) error

	Find(c context.Context, id PhoneNumberID) (*PhoneNumberDTO, error)
	List(c context.Context, q *PhoneNumberQuery) ([]*PhoneNumberDTO, error)
}
