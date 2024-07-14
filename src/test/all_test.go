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
	}

	units.Run(&units.Config{
		Cases:      cname,
		Module:     security.ModuleForTest(),
		Properties: props,
		T:          t,
		UsePanic:   false,
	})

}

func TestUUID(t *testing.T) {
	runTestWithCaseName("test-uuid", t)
}

func TestSubjects(t *testing.T) {
	runTestWithCaseName("test-subjects", t)
}
