package security

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/security/gen/gen4security"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/security"
	theModuleVersion  = "v1.0.14"
	theModuleRevision = 16
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块[github.com/starter-go/security]
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4security.ExportComSetForSecurity)

	mb.Depend(starter.Module())

	return mb.Create()
}
