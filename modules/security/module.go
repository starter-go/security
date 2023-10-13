package security

import (
	"github.com/starter-go/application"
	"github.com/starter-go/security"
	"github.com/starter-go/security/gen/gen4security"
)

// Module ... 导出模块
func Module() application.Module {
	mb := security.ModuleT()
	mb.Components(gen4security.ExportComSetForSecurity)
	return mb.Create()
}
