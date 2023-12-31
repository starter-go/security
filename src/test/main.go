package main

import (
	"os"

	"github.com/starter-go/application"

	"github.com/starter-go/security/modules/security"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(module())
	i.WithPanic(true).Run()
}

////////////////////////////////////////////////////////////////////////////////

func module() application.Module {
	return security.ModuleForTest()
}
