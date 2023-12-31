package main4security

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComSetForSecurity ...
func ExportComSetForSecurity(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
