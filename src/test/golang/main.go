package main

import (
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/units"

	"github.com/starter-go/security/modules/security"
)

func main() {

	a := os.Args
	m := module()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)

}

////////////////////////////////////////////////////////////////////////////////

func module() application.Module {
	return security.ModuleForTest()
}
