package code

import (
	"github.com/starter-go/application"
	"github.com/starter-go/security/keys"
	"github.com/starter-go/vlog"
)

// TestRSA ...
type TestRSA struct {

	//starter:component

	KeysMan keys.Manager //starter:inject("#")
	Logger  vlog.Logger  //starter:inject("#")
}

// Life ...
func (inst *TestRSA) Life() *application.Life {
	return &application.Life{
		Order:  99,
		OnLoop: inst.run,
	}
}

func (inst *TestRSA) run() error {
	list, err := inst.KeysMan.Find("Rsa")
	for _, item := range list {
		inst.Logger.Info("algorithm: %s", item.Name())
	}
	return err
}
