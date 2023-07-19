package rbac

import (
	"time"

	"github.com/starter-go/base/lang"
	"gorm.io/gorm"
)

// BaseEntity 是基本的数据库实体
type BaseEntity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UUID lang.UUID

	Owner     UserID
	Creator   UserID
	Committer UserID
}

// BaseDTO 是基本的 DTO
type BaseDTO struct {
	CreatedAt lang.Time `json:"created_at"`
	UpdatedAt lang.Time `json:"updated_at"`
	DeletedAt lang.Time `json:"deleted_at"`

	UUID lang.UUID `json:"uuid"`

	Owner     UserID `json:"owner"`
	Creator   UserID `json:"creator"`
	Committer UserID `json:"committer"`
}
