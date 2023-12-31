package security

import (
	"github.com/starter-go/application"
	"github.com/starter-go/security"
	"github.com/starter-go/security/gen/main4security"
	"github.com/starter-go/security/gen/test4security"
	"github.com/starter-go/starter"
)

// Module 导出模块 [github.com/starter-go/security]
func Module() application.Module {
	mb := security.NewMainModule()
	mb.Components(main4security.ExportComSetForSecurity)
	mb.Depend(starter.Module())
	return mb.Create()
}

// ModuleForTest 导出模块 [github.com/starter-go/security#test]
func ModuleForTest() application.Module {
	mb := security.NewTestModule()
	mb.Components(test4security.ExportComSetForSecurityTest)
	mb.Depend(Module())
	return mb.Create()
}
