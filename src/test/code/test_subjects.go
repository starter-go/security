package code

import (
	"context"
	"encoding/json"

	"github.com/starter-go/base/context2"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security/subjects"
	"github.com/starter-go/units"
)

// TestSubjects ...
type TestSubjects struct {

	//starter:component

	Loader         subjects.Loader     //starter:inject("#")
	SessionService rbac.SessionService //starter:inject("#")
	SubjectService rbac.SubjectService //starter:inject("#")

}

func (inst *TestSubjects) _impl() units.Units {
	return inst
}

// Units ...
func (inst *TestSubjects) Units(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:    "test-subjects",
		Test:    inst.run,
		Enabled: true,
	}
	list = append(list, r1)
	return list
}

func (inst *TestSubjects) run() error {

	ctx := inst.prepareContext()
	loader := inst.Loader

	// ses1 := new(rbac.SessionDTO)
	// ses2, err := inst.SessionService.Insert(ctx, ses1)
	// if err != nil {
	// 	return err
	// }

	err := subjects.Setup(ctx, loader)
	if err != nil {
		return err
	}

	sub, err := subjects.Current(ctx)
	if err != nil {
		return err
	}

	token := sub.GetToken()

	session := sub.GetSession()
	session.Create()

	// token.SetSessionID(ses2.ID)

	token.SetProperty("t.f1", "a")
	session.SetProperty("s.f1", "b")
	token.SetSessionID(session.SessionID())

	err = session.Commit()
	if err != nil {
		return err
	}

	err = token.Commit()
	if err != nil {
		return err
	}

	return inst.doTestRbacSubjectService(ctx)
}

func (inst *TestSubjects) doTestRbacSubjectService(c context.Context) error {

	item, err := inst.SubjectService.GetCurrent(c)
	if err != nil {
		return err
	}

	json.Marshal(item)

	return nil
}

func (inst *TestSubjects) prepareContext() context.Context {
	ctx := context.Background()
	values := context2.SetupContext(ctx)
	context2.Setup(values)
	return values.Context()
}
