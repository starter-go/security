package test4security

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComSetForSecurityTest ...
func ExportComSetForSecurityTest(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
