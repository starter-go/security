package security

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/starter-go/security"
	theModuleVersion  = "v1.0.54.2"
	theModuleRevision = 62
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule 导出模块[github.com/starter-go/security]
func NewMainModule() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	return mb
}

// NewTestModule 导出模块[github.com/starter-go/security#test]
func NewTestModule() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}
