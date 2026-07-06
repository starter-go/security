package test4security
import (
    p24287f458 "github.com/starter-go/rbac"
    pd4e0ee677 "github.com/starter-go/security"
    p91f218d46 "github.com/starter-go/security/jwt"
    p9621e8b71 "github.com/starter-go/security/random"
    p319728480 "github.com/starter-go/security/src/test/golang/testcom"
    p55f0853be "github.com/starter-go/vlog"
     "github.com/starter-go/application"
)

// type p319728480.MockAuth in package:github.com/starter-go/security/src/test/golang/testcom
//
// id:com-319728480105d3ba-testcom-MockAuth
// class:class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p3197284801_testcom_MockAuth struct {
}

func (inst* p3197284801_testcom_MockAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-319728480105d3ba-testcom-MockAuth"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3197284801_testcom_MockAuth) new() any {
    return &p319728480.MockAuth{}
}

func (inst* p3197284801_testcom_MockAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p319728480.MockAuth)
	nop(ie, com)

	


    return nil
}



// type p319728480.MockTokenAdapter in package:github.com/starter-go/security/src/test/golang/testcom
//
// id:com-319728480105d3ba-testcom-MockTokenAdapter
// class:class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type p3197284801_testcom_MockTokenAdapter struct {
}

func (inst* p3197284801_testcom_MockTokenAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-319728480105d3ba-testcom-MockTokenAdapter"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3197284801_testcom_MockTokenAdapter) new() any {
    return &p319728480.MockTokenAdapter{}
}

func (inst* p3197284801_testcom_MockTokenAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p319728480.MockTokenAdapter)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.Logger = inst.getLogger(ie)


    return nil
}


func (inst*p3197284801_testcom_MockTokenAdapter) getService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p3197284801_testcom_MockTokenAdapter) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}



// type p319728480.MockPermRegDemo in package:github.com/starter-go/security/src/test/golang/testcom
//
// id:com-319728480105d3ba-testcom-MockPermRegDemo
// class:class-08935700f7002f152b848e80701dde49-Registry
// alias:
// scope:singleton
//
type p3197284801_testcom_MockPermRegDemo struct {
}

func (inst* p3197284801_testcom_MockPermRegDemo) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-319728480105d3ba-testcom-MockPermRegDemo"
	r.Classes = "class-08935700f7002f152b848e80701dde49-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3197284801_testcom_MockPermRegDemo) new() any {
    return &p319728480.MockPermRegDemo{}
}

func (inst* p3197284801_testcom_MockPermRegDemo) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p319728480.MockPermRegDemo)
	nop(ie, com)

	


    return nil
}



// type p319728480.TestCom in package:github.com/starter-go/security/src/test/golang/testcom
//
// id:com-319728480105d3ba-testcom-TestCom
// class:
// alias:
// scope:singleton
//
type p3197284801_testcom_TestCom struct {
}

func (inst* p3197284801_testcom_TestCom) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-319728480105d3ba-testcom-TestCom"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3197284801_testcom_TestCom) new() any {
    return &p319728480.TestCom{}
}

func (inst* p3197284801_testcom_TestCom) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p319728480.TestCom)
	nop(ie, com)

	
    com.AuthSer = inst.getAuthSer(ie)
    com.SessionSer = inst.getSessionSer(ie)


    return nil
}


func (inst*p3197284801_testcom_TestCom) getAuthSer(ie application.InjectionExt)p24287f458.AuthService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-AuthService").(p24287f458.AuthService)
}


func (inst*p3197284801_testcom_TestCom) getSessionSer(ie application.InjectionExt)pd4e0ee677.SessionService{
    return ie.GetComponent("#alias-d4e0ee677c339b7ffcf1d55767953499-SessionService").(pd4e0ee677.SessionService)
}



// type p319728480.TestRandom in package:github.com/starter-go/security/src/test/golang/testcom
//
// id:com-319728480105d3ba-testcom-TestRandom
// class:
// alias:
// scope:singleton
//
type p3197284801_testcom_TestRandom struct {
}

func (inst* p3197284801_testcom_TestRandom) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-319728480105d3ba-testcom-TestRandom"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3197284801_testcom_TestRandom) new() any {
    return &p319728480.TestRandom{}
}

func (inst* p3197284801_testcom_TestRandom) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p319728480.TestRandom)
	nop(ie, com)

	
    com.Rand = inst.getRand(ie)
    com.Logger = inst.getLogger(ie)


    return nil
}


func (inst*p3197284801_testcom_TestRandom) getRand(ie application.InjectionExt)p9621e8b71.Service{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-Service").(p9621e8b71.Service)
}


func (inst*p3197284801_testcom_TestRandom) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}



// type p319728480.TestUUID in package:github.com/starter-go/security/src/test/golang/testcom
//
// id:com-319728480105d3ba-testcom-TestUUID
// class:
// alias:
// scope:singleton
//
type p3197284801_testcom_TestUUID struct {
}

func (inst* p3197284801_testcom_TestUUID) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-319728480105d3ba-testcom-TestUUID"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3197284801_testcom_TestUUID) new() any {
    return &p319728480.TestUUID{}
}

func (inst* p3197284801_testcom_TestUUID) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p319728480.TestUUID)
	nop(ie, com)

	
    com.Ser = inst.getSer(ie)
    com.Logger = inst.getLogger(ie)


    return nil
}


func (inst*p3197284801_testcom_TestUUID) getSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}


func (inst*p3197284801_testcom_TestUUID) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}


