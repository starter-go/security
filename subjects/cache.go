package subjects

import (
	"context"

	"github.com/starter-go/v0/subjects"
)

// Current 从上下文中提取当前主体
func Current(c context.Context) (Subject, error) {

	return subjects.GetCurrent(c)
}
