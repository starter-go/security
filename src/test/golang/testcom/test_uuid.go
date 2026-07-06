package testcom

import (
	"context"

	"github.com/starter-go/security/random"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// TestUUID ...
type TestUUID struct {

	//starter:component

	Ser    random.UUIDService //starter:inject("#")
	Logger vlog.Logger        //starter:inject("#")

}

func (inst *TestUUID) _impl() units.Unit {
	return inst
}

// Units ...
func (inst *TestUUID) ListRegistrations(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:    "test-uuid",
		Do:      inst.run,
		Enabled: true,
	}
	list = append(list, r1)
	return list
}

func (inst *TestUUID) run(cc context.Context) error {

	u1 := inst.Ser.Build().Generate()
	u2 := inst.Ser.Build().Generate()
	u3 := inst.Ser.Build().Generate()

	inst.Logger.Info("test: uuid-service: gen uuid: %s", u1.String())
	inst.Logger.Info("test: uuid-service: gen uuid: %s", u2.String())
	inst.Logger.Info("test: uuid-service: gen uuid: %s", u3.String())

	return nil
}
