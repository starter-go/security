package rbac

import (
	"time"

	"github.com/starter-go/base/lang"
)

// BaseVO 是通用的基本 VO 结构
type BaseVO struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	Time       time.Time   `json:"time"`
	Timestamp  lang.Time   `json:"timestamp"`
	Pagination *Pagination `json:"pagination"`
}

////////////////////////////////////////////////////////////////////////////////

// VOGetter 是获取 BaseVO 的接口
type VOGetter interface {
	GetVO() *BaseVO
}

////////////////////////////////////////////////////////////////////////////////

// GetVO 实现 BaseGetter
func (inst *BaseVO) GetVO() *BaseVO {
	return inst
}
