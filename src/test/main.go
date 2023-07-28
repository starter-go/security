package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modgorm"
	"github.com/starter-go/security"
	"github.com/starter-go/security/gen/gen4test"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(module())
	i.WithPanic(true).Run()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed "resources"
var theModuleResFS embed.FS

func module() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("security/src/test")
	mb.Version("0.0.1")
	mb.Revision(1)
	mb.EmbedResources(theModuleResFS, "resources")
	mb.Components(gen4test.ExportComSetForSecurityTest)

	mb.Depend(security.Module())
	mb.Depend(modgorm.Module())

	return mb.Create()
}
