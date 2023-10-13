package rbac

import (
	"time"

	"github.com/starter-go/base/lang"
)

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

// BaseVO 是通用的基本 VO 结构
type BaseVO struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	Time       time.Time   `json:"time"`
	Timestamp  lang.Time   `json:"timestamp"`
	Pagination *Pagination `json:"pagination"`
}
