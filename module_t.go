package security

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/security"
	theModuleVersion  = "v1.0.35"
	theModuleRevision = 37
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// ModuleT 导出模块[github.com/starter-go/security]
func ModuleT() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	// mb.Components(gen4security.ExportComSetForSecurity)

	mb.Depend(starter.Module())

	return mb
}
