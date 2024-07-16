package auth

import "github.com/starter-go/application/parameters"

// Feedback ... 用于向请求方反馈验证信息
type Feedback interface {
	Parameters() parameters.Table
}

// NewFeedback 新建一个默认的 Feedback 对象
func NewFeedback() Feedback {
	fb := &defaultFeedback{}
	fb.init()
	return fb
}

////////////////////////////////////////////////////////////////////////////////

type defaultFeedback struct {
	params parameters.Table
}

func (inst *defaultFeedback) _impl() Feedback { return inst }

func (inst *defaultFeedback) init() {
	inst.params = parameters.NewTable(nil)
}

func (inst *defaultFeedback) Parameters() parameters.Table {
	return inst.params
}

////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
