package subjects

import (
	"context"
	"fmt"

	"github.com/starter-go/base/context2"
)

const theSubjectBindingKey = "security.subjects.Subject#binding"

type cache struct {
	c       context.Context
	subject Subject
	loader  Loader
}

func getCache(c context.Context, create bool) (*cache, error) {
	const key = theSubjectBindingKey
	values, err := context2.GetValues(c)
	if err != nil {
		return nil, err
	}
	o1 := values.GetValue(key)
	if o1 == nil {
		if create {
			o2 := new(cache)
			values.SetValue(key, o2)
			return o2, nil
		}
		return nil, fmt.Errorf("call subjects.GetSubject() without setup")
	}
	o2 := o1.(*cache)
	return o2, nil
}

func (inst *cache) fetch() (Subject, error) {
	sub := inst.subject
	if sub != nil {
		return sub, nil
	}
	ctx := inst.c
	sub, err := inst.loader.Load(ctx)
	if err != nil {
		return nil, err
	}
	inst.subject = sub
	return sub, nil
}

////////////////////////////////////////////////////////////////////////////////

// Current 从上下文中提取当前主体
func Current(c context.Context) (Subject, error) {
	ca, err := getCache(c, false)
	if err != nil {
		return nil, err
	}
	return ca.fetch()
}

// Setup 把主体绑定到上下文中
func Setup(c context.Context, loader Loader) error {
	if loader == nil || c == nil {
		return fmt.Errorf("param is nil")
	}
	ca, err := getCache(c, true)
	if err != nil {
		return err
	}
	ca.loader = loader
	return nil
}
