package rbac

import (
	"context"

	"gorm.io/gorm"
)

// PhoneNumberID ...
type PhoneNumberID int64

// PhoneNumberEntity ...
type PhoneNumberEntity struct {
	ID PhoneNumberID

	BaseDTO

	SimpleNunber string `gorm:"unique"`
}

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

// PhoneNumberDAO ...
type PhoneNumberDAO interface {
	Insert(db *gorm.DB, o *PhoneNumberEntity) (*PhoneNumberEntity, error)
	Update(db *gorm.DB, id PhoneNumberID, updater func(*PhoneNumberEntity)) (*PhoneNumberEntity, error)
	Delete(db *gorm.DB, id PhoneNumberID) error

	Find(db *gorm.DB, id PhoneNumberID) (*PhoneNumberEntity, error)
	List(db *gorm.DB, q *PhoneNumberQuery) ([]*PhoneNumberEntity, error)
}

// PhoneNumberService ...
type PhoneNumberService interface {
	Insert(c context.Context, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Update(c context.Context, id PhoneNumberID, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Delete(c context.Context, id PhoneNumberID) error

	Find(c context.Context, id PhoneNumberID) (*PhoneNumberDTO, error)
	List(c context.Context, q *PhoneNumberQuery) ([]*PhoneNumberDTO, error)
}

// PhoneNumberConvertor ...
type PhoneNumberConvertor interface {
	ConvertE2D(c context.Context, entity *PhoneNumberEntity) (*PhoneNumberDTO, error)
	ConvertD2E(c context.Context, dto *PhoneNumberDTO) (*PhoneNumberEntity, error)

	ConvertListE2D(c context.Context, entity []*PhoneNumberEntity) ([]*PhoneNumberDTO, error)
}
