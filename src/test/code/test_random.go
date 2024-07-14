package code

import (
	"github.com/starter-go/security/random"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// TestRandom ...
type TestRandom struct {

	//starter:component

	Rand   random.Service //starter:inject("#")
	Logger vlog.Logger    //starter:inject("#")

}

func (inst *TestRandom) _impl() units.Units {
	return inst
}

// Units ...
func (inst *TestRandom) Units(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:    "test-random",
		Test:    inst.run,
		Enabled: true,
	}
	list = append(list, r1)
	return list
}

func (inst *TestRandom) run() error {

	r := inst.Rand
	l := inst.Logger

	n := r.NextInt()
	n64 := r.NextInt64()
	str := r.NextString(20)
	data := r.NextBytes(20)

	l.Warn("n=%d \nn64=%d \nstr=%s \ndata=%s", n, n64, str, data)

	return nil
}
