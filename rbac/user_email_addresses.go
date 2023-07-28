package rbac

import (
	"context"

	"gorm.io/gorm"
)

// EmailAddressID ...
type EmailAddressID int64

// EmailAddressEntity ...
type EmailAddressEntity struct {
	ID EmailAddressID

	BaseEntity

	Address string `gorm:"unique"`
}

// EmailAddressDTO ...
type EmailAddressDTO struct {
	ID EmailAddressID `json:"id"`

	BaseDTO

	Address string `json:"address"`
}

// EmailAddressQuery 查询参数
type EmailAddressQuery struct {
	Pagination Pagination
}

// EmailAddressDAO ...
type EmailAddressDAO interface {
	Insert(db *gorm.DB, o *EmailAddressEntity) (*EmailAddressEntity, error)
	Update(db *gorm.DB, id EmailAddressID, updater func(*EmailAddressEntity)) (*EmailAddressEntity, error)
	Delete(db *gorm.DB, id EmailAddressID) error

	Find(db *gorm.DB, id EmailAddressID) (*EmailAddressEntity, error)
	List(db *gorm.DB, q *EmailAddressQuery) ([]*EmailAddressEntity, error)
}

// EmailAddressService ...
type EmailAddressService interface {
	Insert(c context.Context, o *EmailAddressDTO) (*EmailAddressDTO, error)
	Update(c context.Context, id EmailAddressID, o *EmailAddressDTO) (*EmailAddressDTO, error)
	Delete(c context.Context, id EmailAddressID) error

	Find(c context.Context, id EmailAddressID) (*EmailAddressDTO, error)
	List(c context.Context, q *EmailAddressQuery) ([]*EmailAddressDTO, error)
}

// EmailAddressConvertor ...
type EmailAddressConvertor interface {
	ConvertE2D(c context.Context, entity *EmailAddressEntity) (*EmailAddressDTO, error)
	ConvertD2E(c context.Context, dto *EmailAddressDTO) (*EmailAddressEntity, error)

	ConvertListE2D(c context.Context, entity []*EmailAddressEntity) ([]*EmailAddressDTO, error)
}
