package test4security
import (
    p24287f458 "github.com/starter-go/rbac"
    pd4e0ee677 "github.com/starter-go/security"
    p91f218d46 "github.com/starter-go/security/jwt"
    p9621e8b71 "github.com/starter-go/security/random"
    p577233893 "github.com/starter-go/security/src/test/code"
    paff1180b7 "github.com/starter-go/security/subjects"
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


    return nil
}


func (inst*p5772338936_code_TestCom) getAuthSer(ie application.InjectionExt)p24287f458.AuthService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-AuthService").(p24287f458.AuthService)
}


func (inst*p5772338936_code_TestCom) getSessionSer(ie application.InjectionExt)pd4e0ee677.SessionService{
    return ie.GetComponent("#alias-d4e0ee677c339b7ffcf1d55767953499-SessionService").(pd4e0ee677.SessionService)
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



// type p577233893.TestSubjects in package:github.com/starter-go/security/src/test/code
//
// id:com-5772338936071413-code-TestSubjects
// class:
// alias:
// scope:singleton
//
type p5772338936_code_TestSubjects struct {
}

func (inst* p5772338936_code_TestSubjects) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5772338936071413-code-TestSubjects"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5772338936_code_TestSubjects) new() any {
    return &p577233893.TestSubjects{}
}

func (inst* p5772338936_code_TestSubjects) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p577233893.TestSubjects)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)
    com.SessionService = inst.getSessionService(ie)


    return nil
}


func (inst*p5772338936_code_TestSubjects) getLoader(ie application.InjectionExt)paff1180b7.Loader{
    return ie.GetComponent("#alias-aff1180b734cd089659a2dcc3be458d7-Loader").(paff1180b7.Loader)
}


func (inst*p5772338936_code_TestSubjects) getSessionService(ie application.InjectionExt)p24287f458.SessionService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-SessionService").(p24287f458.SessionService)
}



// type p577233893.TestUUID in package:github.com/starter-go/security/src/test/code
//
// id:com-5772338936071413-code-TestUUID
// class:
// alias:
// scope:singleton
//
type p5772338936_code_TestUUID struct {
}

func (inst* p5772338936_code_TestUUID) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5772338936071413-code-TestUUID"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5772338936_code_TestUUID) new() any {
    return &p577233893.TestUUID{}
}

func (inst* p5772338936_code_TestUUID) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p577233893.TestUUID)
	nop(ie, com)

	
    com.Ser = inst.getSer(ie)
    com.Logger = inst.getLogger(ie)


    return nil
}


func (inst*p5772338936_code_TestUUID) getSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}


func (inst*p5772338936_code_TestUUID) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}


