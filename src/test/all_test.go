package main

import (
	"testing"

	"github.com/starter-go/security/modules/security"
	"github.com/starter-go/units"
)

func runTestWithCaseName(cname string, t *testing.T) {

	props := map[string]string{
		"debug.enabled":        "1",
		"debug.log-properties": "1",

		"test.case.cname": cname,
	}

	ctx := units.NewContext()
	m := security.ModuleForTest()

	ctx.T = t
	ctx.Module = m
	ctx.Properties = props
	ctx.UsePanic = true

	units.Run(ctx)

}

func TestUUID(t *testing.T) {
	runTestWithCaseName("test-uuid", t)
}

func TestSubjects(t *testing.T) {
	runTestWithCaseName("test-subjects", t)
}
