package rbac

import (
	"time"

	"github.com/starter-go/base/lang"
	"gorm.io/gorm"
)

// BaseEntity 是基本的数据库实体
type BaseEntity struct {
	UUID lang.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Owner   UserID
	Creator UserID
	Updater UserID
}

// BaseDTO 是基本的 DTO
type BaseDTO struct {
	UUID lang.UUID `json:"uuid"`

	CreatedAt lang.Time `json:"created_at"`
	UpdatedAt lang.Time `json:"updated_at"`
	DeletedAt lang.Time `json:"deleted_at"`

	Owner   UserID `json:"owner"`
	Creator UserID `json:"creator"`
	Updater UserID `json:"updater"`
}
