package gen4test
import (
    p91f218d46 "github.com/starter-go/security/jwt"
    p9621e8b71 "github.com/starter-go/security/random"
    p2dece1e49 "github.com/starter-go/security/rbac"
    p577233893 "github.com/starter-go/security/src/test/code"
    p55f0853be "github.com/starter-go/vlog"
     "github.com/starter-go/application"
)

// type p577233893.MockAuth in package:github.com/starter-go/security/src/test/code
//
// id:com-5772338936071413-code-MockAuth
// class:class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p5772338936_code_MockAuth struct {
}

func (inst* p5772338936_code_MockAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5772338936071413-code-MockAuth"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5772338936_code_MockAuth) new() any {
    return &p577233893.MockAuth{}
}

func (inst* p5772338936_code_MockAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p577233893.MockAuth)
	nop(ie, com)

	


    return nil
}



// type p577233893.MockTokenAdapter in package:github.com/starter-go/security/src/test/code
//
// id:com-5772338936071413-code-MockTokenAdapter
// class:class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type p5772338936_code_MockTokenAdapter struct {
}

func (inst* p5772338936_code_MockTokenAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5772338936071413-code-MockTokenAdapter"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5772338936_code_MockTokenAdapter) new() any {
    return &p577233893.MockTokenAdapter{}
}

func (inst* p5772338936_code_MockTokenAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p577233893.MockTokenAdapter)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.Logger = inst.getLogger(ie)


    return nil
}


func (inst*p5772338936_code_MockTokenAdapter) getService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p5772338936_code_MockTokenAdapter) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}



// type p577233893.TestCom in package:github.com/starter-go/security/src/test/code
//
// id:com-5772338936071413-code-TestCom
// class:
// alias:
// scope:singleton
//
type p5772338936_code_TestCom struct {
}

func (inst* p5772338936_code_TestCom) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5772338936071413-code-TestCom"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5772338936_code_TestCom) new() any {
    return &p577233893.TestCom{}
}

func (inst* p5772338936_code_TestCom) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p577233893.TestCom)
	nop(ie, com)

	
    com.AuthSer = inst.getAuthSer(ie)
    com.SessionSer = inst.getSessionSer(ie)
    com.UserSer = inst.getUserSer(ie)
    com.PermissionSer = inst.getPermissionSer(ie)


    return nil
}


func (inst*p5772338936_code_TestCom) getAuthSer(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}


func (inst*p5772338936_code_TestCom) getSessionSer(ie application.InjectionExt)p2dece1e49.SessionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-SessionService").(p2dece1e49.SessionService)
}


func (inst*p5772338936_code_TestCom) getUserSer(ie application.InjectionExt)p2dece1e49.UserService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}


func (inst*p5772338936_code_TestCom) getPermissionSer(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p577233893.TestRandom in package:github.com/starter-go/security/src/test/code
//
// id:com-5772338936071413-code-TestRandom
// class:
// alias:
// scope:singleton
//
type p5772338936_code_TestRandom struct {
}

func (inst* p5772338936_code_TestRandom) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5772338936071413-code-TestRandom"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5772338936_code_TestRandom) new() any {
    return &p577233893.TestRandom{}
}

func (inst* p5772338936_code_TestRandom) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p577233893.TestRandom)
	nop(ie, com)

	
    com.Rand = inst.getRand(ie)
    com.Logger = inst.getLogger(ie)


    return nil
}


func (inst*p5772338936_code_TestRandom) getRand(ie application.InjectionExt)p9621e8b71.Service{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-Service").(p9621e8b71.Service)
}


func (inst*p5772338936_code_TestRandom) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}


