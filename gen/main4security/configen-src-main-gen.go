package main4security
import (
    pd4e0ee677 "github.com/starter-go/security"
    p9d209f7c2 "github.com/starter-go/security/auth"
    pf41d62225 "github.com/starter-go/security/internal"
    p91f218d46 "github.com/starter-go/security/jwt"
    p9621e8b71 "github.com/starter-go/security/random"
     "github.com/starter-go/application"
)

// type pf41d62225.AuthService1 in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-AuthService1
// class:
// alias:alias-9d209f7c2504d33e6054a2c9998e9485-Service
// scope:singleton
//
type pf41d62225c_internal_AuthService1 struct {
}

func (inst* pf41d62225c_internal_AuthService1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-AuthService1"
	r.Classes = ""
	r.Aliases = "alias-9d209f7c2504d33e6054a2c9998e9485-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_AuthService1) new() any {
    return &pf41d62225.AuthService1{}
}

func (inst* pf41d62225c_internal_AuthService1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.AuthService1)
	nop(ie, com)

	
    com.Authenticators = inst.getAuthenticators(ie)
    com.Authorizers = inst.getAuthorizers(ie)
    com.LogBadAuthError = inst.getLogBadAuthError(ie)


    return nil
}


func (inst*pf41d62225c_internal_AuthService1) getAuthenticators(ie application.InjectionExt)p9d209f7c2.AuthenticatorManager{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-AuthenticatorManager").(p9d209f7c2.AuthenticatorManager)
}


func (inst*pf41d62225c_internal_AuthService1) getAuthorizers(ie application.InjectionExt)p9d209f7c2.AuthorizerManager{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-AuthorizerManager").(p9d209f7c2.AuthorizerManager)
}


func (inst*pf41d62225c_internal_AuthService1) getLogBadAuthError(ie application.InjectionExt)bool{
    return ie.GetBool("${security.log-bad-auth-error}")
}



// type pf41d62225.AuthService2 in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-AuthService2
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-AuthService
// scope:singleton
//
type pf41d62225c_internal_AuthService2 struct {
}

func (inst* pf41d62225c_internal_AuthService2) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-AuthService2"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-AuthService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_AuthService2) new() any {
    return &pf41d62225.AuthService2{}
}

func (inst* pf41d62225c_internal_AuthService2) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.AuthService2)
	nop(ie, com)

	
    com.Servic1 = inst.getServic1(ie)


    return nil
}


func (inst*pf41d62225c_internal_AuthService2) getServic1(ie application.InjectionExt)p9d209f7c2.Service{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-Service").(p9d209f7c2.Service)
}



// type pf41d62225.AuthenticatorManagerImpl in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-AuthenticatorManagerImpl
// class:
// alias:alias-9d209f7c2504d33e6054a2c9998e9485-AuthenticatorManager
// scope:singleton
//
type pf41d62225c_internal_AuthenticatorManagerImpl struct {
}

func (inst* pf41d62225c_internal_AuthenticatorManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-AuthenticatorManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-9d209f7c2504d33e6054a2c9998e9485-AuthenticatorManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_AuthenticatorManagerImpl) new() any {
    return &pf41d62225.AuthenticatorManagerImpl{}
}

func (inst* pf41d62225c_internal_AuthenticatorManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.AuthenticatorManagerImpl)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*pf41d62225c_internal_AuthenticatorManagerImpl) getRegistryList(ie application.InjectionExt)[]p9d209f7c2.Registry{
    dst := make([]p9d209f7c2.Registry, 0)
    src := ie.ListComponents(".class-9d209f7c2504d33e6054a2c9998e9485-Registry")
    for _, item1 := range src {
        item2 := item1.(p9d209f7c2.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type pf41d62225.AuthorizerManagerImpl in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-AuthorizerManagerImpl
// class:
// alias:alias-9d209f7c2504d33e6054a2c9998e9485-AuthorizerManager
// scope:singleton
//
type pf41d62225c_internal_AuthorizerManagerImpl struct {
}

func (inst* pf41d62225c_internal_AuthorizerManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-AuthorizerManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-9d209f7c2504d33e6054a2c9998e9485-AuthorizerManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_AuthorizerManagerImpl) new() any {
    return &pf41d62225.AuthorizerManagerImpl{}
}

func (inst* pf41d62225c_internal_AuthorizerManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.AuthorizerManagerImpl)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*pf41d62225c_internal_AuthorizerManagerImpl) getRegistryList(ie application.InjectionExt)[]p9d209f7c2.Registry{
    dst := make([]p9d209f7c2.Registry, 0)
    src := ie.ListComponents(".class-9d209f7c2504d33e6054a2c9998e9485-Registry")
    for _, item1 := range src {
        item2 := item1.(p9d209f7c2.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type pf41d62225.DefaultRandomService in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-DefaultRandomService
// class:
// alias:alias-9621e8b71013b0fc25942a1749ed3652-Service
// scope:singleton
//
type pf41d62225c_internal_DefaultRandomService struct {
}

func (inst* pf41d62225c_internal_DefaultRandomService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-DefaultRandomService"
	r.Classes = ""
	r.Aliases = "alias-9621e8b71013b0fc25942a1749ed3652-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_DefaultRandomService) new() any {
    return &pf41d62225.DefaultRandomService{}
}

func (inst* pf41d62225c_internal_DefaultRandomService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.DefaultRandomService)
	nop(ie, com)

	
    com.Providers = inst.getProviders(ie)


    return nil
}


func (inst*pf41d62225c_internal_DefaultRandomService) getProviders(ie application.InjectionExt)[]p9621e8b71.ProviderRegistry{
    dst := make([]p9621e8b71.ProviderRegistry, 0)
    src := ie.ListComponents(".class-9621e8b71013b0fc25942a1749ed3652-ProviderRegistry")
    for _, item1 := range src {
        item2 := item1.(p9621e8b71.ProviderRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type pf41d62225.DefaultRandomSource in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-DefaultRandomSource
// class:class-9621e8b71013b0fc25942a1749ed3652-ProviderRegistry
// alias:
// scope:singleton
//
type pf41d62225c_internal_DefaultRandomSource struct {
}

func (inst* pf41d62225c_internal_DefaultRandomSource) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-DefaultRandomSource"
	r.Classes = "class-9621e8b71013b0fc25942a1749ed3652-ProviderRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_DefaultRandomSource) new() any {
    return &pf41d62225.DefaultRandomSource{}
}

func (inst* pf41d62225c_internal_DefaultRandomSource) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.DefaultRandomSource)
	nop(ie, com)

	


    return nil
}



// type pf41d62225.JWTCODEC in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-JWTCODEC
// class:class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type pf41d62225c_internal_JWTCODEC struct {
}

func (inst* pf41d62225c_internal_JWTCODEC) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-JWTCODEC"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_JWTCODEC) new() any {
    return &pf41d62225.JWTCODEC{}
}

func (inst* pf41d62225c_internal_JWTCODEC) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.JWTCODEC)
	nop(ie, com)

	


    return nil
}



// type pf41d62225.JWTService in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-JWTService
// class:
// alias:alias-91f218d46ec21cd234778bbe54aecc66-Service
// scope:singleton
//
type pf41d62225c_internal_JWTService struct {
}

func (inst* pf41d62225c_internal_JWTService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-JWTService"
	r.Classes = ""
	r.Aliases = "alias-91f218d46ec21cd234778bbe54aecc66-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_JWTService) new() any {
    return &pf41d62225.JWTService{}
}

func (inst* pf41d62225c_internal_JWTService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.JWTService)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)
    com.TokenMaxageDefault = inst.getTokenMaxageDefault(ie)
    com.TokenMaxageMin = inst.getTokenMaxageMin(ie)
    com.TokenMaxageMax = inst.getTokenMaxageMax(ie)
    com.SessionMaxageDefault = inst.getSessionMaxageDefault(ie)
    com.SessionMaxageMin = inst.getSessionMaxageMin(ie)
    com.SessionMaxageMax = inst.getSessionMaxageMax(ie)


    return nil
}


func (inst*pf41d62225c_internal_JWTService) getRegistryList(ie application.InjectionExt)[]p91f218d46.Registry{
    dst := make([]p91f218d46.Registry, 0)
    src := ie.ListComponents(".class-91f218d46ec21cd234778bbe54aecc66-Registry")
    for _, item1 := range src {
        item2 := item1.(p91f218d46.Registry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pf41d62225c_internal_JWTService) getTokenMaxageDefault(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.default-age-sec}")
}


func (inst*pf41d62225c_internal_JWTService) getTokenMaxageMin(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.min-age-sec}")
}


func (inst*pf41d62225c_internal_JWTService) getTokenMaxageMax(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.max-age-sec}")
}


func (inst*pf41d62225c_internal_JWTService) getSessionMaxageDefault(ie application.InjectionExt)int{
    return ie.GetInt("${security.session.default-age-sec}")
}


func (inst*pf41d62225c_internal_JWTService) getSessionMaxageMin(ie application.InjectionExt)int{
    return ie.GetInt("${security.session.min-age-sec}")
}


func (inst*pf41d62225c_internal_JWTService) getSessionMaxageMax(ie application.InjectionExt)int{
    return ie.GetInt("${security.session.max-age-sec}")
}



// type pf41d62225.SessionServiceImpl in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-SessionServiceImpl
// class:
// alias:alias-d4e0ee677c339b7ffcf1d55767953499-SessionService
// scope:singleton
//
type pf41d62225c_internal_SessionServiceImpl struct {
}

func (inst* pf41d62225c_internal_SessionServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-SessionServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-d4e0ee677c339b7ffcf1d55767953499-SessionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_SessionServiceImpl) new() any {
    return &pf41d62225.SessionServiceImpl{}
}

func (inst* pf41d62225c_internal_SessionServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.SessionServiceImpl)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*pf41d62225c_internal_SessionServiceImpl) getRegistryList(ie application.InjectionExt)[]pd4e0ee677.SessionRegistry{
    dst := make([]pd4e0ee677.SessionRegistry, 0)
    src := ie.ListComponents(".class-d4e0ee677c339b7ffcf1d55767953499-SessionRegistry")
    for _, item1 := range src {
        item2 := item1.(pd4e0ee677.SessionRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type pf41d62225.SubjectServiceImpl in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-SubjectServiceImpl
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-SubjectService
// scope:singleton
//
type pf41d62225c_internal_SubjectServiceImpl struct {
}

func (inst* pf41d62225c_internal_SubjectServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-SubjectServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-SubjectService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_SubjectServiceImpl) new() any {
    return &pf41d62225.SubjectServiceImpl{}
}

func (inst* pf41d62225c_internal_SubjectServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.SubjectServiceImpl)
	nop(ie, com)

	
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*pf41d62225c_internal_SubjectServiceImpl) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type pf41d62225.UUIDServiceImpl in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-UUIDServiceImpl
// class:
// alias:alias-9621e8b71013b0fc25942a1749ed3652-UUIDService
// scope:singleton
//
type pf41d62225c_internal_UUIDServiceImpl struct {
}

func (inst* pf41d62225c_internal_UUIDServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-UUIDServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-9621e8b71013b0fc25942a1749ed3652-UUIDService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf41d62225c_internal_UUIDServiceImpl) new() any {
    return &pf41d62225.UUIDServiceImpl{}
}

func (inst* pf41d62225c_internal_UUIDServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf41d62225.UUIDServiceImpl)
	nop(ie, com)

	
    com.Host = inst.getHost(ie)


    return nil
}


func (inst*pf41d62225c_internal_UUIDServiceImpl) getHost(ie application.InjectionExt)string{
    return ie.GetString("${security.uuid.service.hostname}")
}


