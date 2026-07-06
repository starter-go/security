package main4security
import (
    pd4e0ee677 "github.com/starter-go/security"
    p9d209f7c2 "github.com/starter-go/security/auth"
    p1579242dd "github.com/starter-go/security/impl"
    pbf7e6103c "github.com/starter-go/security/impl/ipermissions"
    p91f218d46 "github.com/starter-go/security/jwt"
    p08935700f "github.com/starter-go/security/permissions"
    p9621e8b71 "github.com/starter-go/security/random"
     "github.com/starter-go/application"
)

// type p1579242dd.AuthService1 in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-AuthService1
// class:
// alias:alias-9d209f7c2504d33e6054a2c9998e9485-Service
// scope:singleton
//
type p1579242dd0_impl_AuthService1 struct {
}

func (inst* p1579242dd0_impl_AuthService1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-AuthService1"
	r.Classes = ""
	r.Aliases = "alias-9d209f7c2504d33e6054a2c9998e9485-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_AuthService1) new() any {
    return &p1579242dd.AuthService1{}
}

func (inst* p1579242dd0_impl_AuthService1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.AuthService1)
	nop(ie, com)

	
    com.Authenticators = inst.getAuthenticators(ie)
    com.Authorizers = inst.getAuthorizers(ie)
    com.LogBadAuthError = inst.getLogBadAuthError(ie)


    return nil
}


func (inst*p1579242dd0_impl_AuthService1) getAuthenticators(ie application.InjectionExt)p9d209f7c2.AuthenticatorManager{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-AuthenticatorManager").(p9d209f7c2.AuthenticatorManager)
}


func (inst*p1579242dd0_impl_AuthService1) getAuthorizers(ie application.InjectionExt)p9d209f7c2.AuthorizerManager{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-AuthorizerManager").(p9d209f7c2.AuthorizerManager)
}


func (inst*p1579242dd0_impl_AuthService1) getLogBadAuthError(ie application.InjectionExt)bool{
    return ie.GetBool("${security.log-bad-auth-error}")
}



// type p1579242dd.AuthService2 in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-AuthService2
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-AuthService
// scope:singleton
//
type p1579242dd0_impl_AuthService2 struct {
}

func (inst* p1579242dd0_impl_AuthService2) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-AuthService2"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-AuthService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_AuthService2) new() any {
    return &p1579242dd.AuthService2{}
}

func (inst* p1579242dd0_impl_AuthService2) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.AuthService2)
	nop(ie, com)

	
    com.Servic1 = inst.getServic1(ie)


    return nil
}


func (inst*p1579242dd0_impl_AuthService2) getServic1(ie application.InjectionExt)p9d209f7c2.Service{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-Service").(p9d209f7c2.Service)
}



// type p1579242dd.AuthenticatorManagerImpl in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-AuthenticatorManagerImpl
// class:
// alias:alias-9d209f7c2504d33e6054a2c9998e9485-AuthenticatorManager
// scope:singleton
//
type p1579242dd0_impl_AuthenticatorManagerImpl struct {
}

func (inst* p1579242dd0_impl_AuthenticatorManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-AuthenticatorManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-9d209f7c2504d33e6054a2c9998e9485-AuthenticatorManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_AuthenticatorManagerImpl) new() any {
    return &p1579242dd.AuthenticatorManagerImpl{}
}

func (inst* p1579242dd0_impl_AuthenticatorManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.AuthenticatorManagerImpl)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*p1579242dd0_impl_AuthenticatorManagerImpl) getRegistryList(ie application.InjectionExt)[]p9d209f7c2.Registry{
    dst := make([]p9d209f7c2.Registry, 0)
    src := ie.ListComponents(".class-9d209f7c2504d33e6054a2c9998e9485-Registry")
    for _, item1 := range src {
        item2 := item1.(p9d209f7c2.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type p1579242dd.AuthorizerManagerImpl in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-AuthorizerManagerImpl
// class:
// alias:alias-9d209f7c2504d33e6054a2c9998e9485-AuthorizerManager
// scope:singleton
//
type p1579242dd0_impl_AuthorizerManagerImpl struct {
}

func (inst* p1579242dd0_impl_AuthorizerManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-AuthorizerManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-9d209f7c2504d33e6054a2c9998e9485-AuthorizerManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_AuthorizerManagerImpl) new() any {
    return &p1579242dd.AuthorizerManagerImpl{}
}

func (inst* p1579242dd0_impl_AuthorizerManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.AuthorizerManagerImpl)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*p1579242dd0_impl_AuthorizerManagerImpl) getRegistryList(ie application.InjectionExt)[]p9d209f7c2.Registry{
    dst := make([]p9d209f7c2.Registry, 0)
    src := ie.ListComponents(".class-9d209f7c2504d33e6054a2c9998e9485-Registry")
    for _, item1 := range src {
        item2 := item1.(p9d209f7c2.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type p1579242dd.DefaultRandomService in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-DefaultRandomService
// class:
// alias:alias-9621e8b71013b0fc25942a1749ed3652-Service
// scope:singleton
//
type p1579242dd0_impl_DefaultRandomService struct {
}

func (inst* p1579242dd0_impl_DefaultRandomService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-DefaultRandomService"
	r.Classes = ""
	r.Aliases = "alias-9621e8b71013b0fc25942a1749ed3652-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_DefaultRandomService) new() any {
    return &p1579242dd.DefaultRandomService{}
}

func (inst* p1579242dd0_impl_DefaultRandomService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.DefaultRandomService)
	nop(ie, com)

	
    com.Providers = inst.getProviders(ie)


    return nil
}


func (inst*p1579242dd0_impl_DefaultRandomService) getProviders(ie application.InjectionExt)[]p9621e8b71.ProviderRegistry{
    dst := make([]p9621e8b71.ProviderRegistry, 0)
    src := ie.ListComponents(".class-9621e8b71013b0fc25942a1749ed3652-ProviderRegistry")
    for _, item1 := range src {
        item2 := item1.(p9621e8b71.ProviderRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p1579242dd.DefaultRandomSource in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-DefaultRandomSource
// class:class-9621e8b71013b0fc25942a1749ed3652-ProviderRegistry
// alias:
// scope:singleton
//
type p1579242dd0_impl_DefaultRandomSource struct {
}

func (inst* p1579242dd0_impl_DefaultRandomSource) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-DefaultRandomSource"
	r.Classes = "class-9621e8b71013b0fc25942a1749ed3652-ProviderRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_DefaultRandomSource) new() any {
    return &p1579242dd.DefaultRandomSource{}
}

func (inst* p1579242dd0_impl_DefaultRandomSource) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.DefaultRandomSource)
	nop(ie, com)

	


    return nil
}



// type p1579242dd.JWTCODEC in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-JWTCODEC
// class:class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type p1579242dd0_impl_JWTCODEC struct {
}

func (inst* p1579242dd0_impl_JWTCODEC) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-JWTCODEC"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_JWTCODEC) new() any {
    return &p1579242dd.JWTCODEC{}
}

func (inst* p1579242dd0_impl_JWTCODEC) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.JWTCODEC)
	nop(ie, com)

	


    return nil
}



// type p1579242dd.JWTService in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-JWTService
// class:
// alias:alias-91f218d46ec21cd234778bbe54aecc66-Service
// scope:singleton
//
type p1579242dd0_impl_JWTService struct {
}

func (inst* p1579242dd0_impl_JWTService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-JWTService"
	r.Classes = ""
	r.Aliases = "alias-91f218d46ec21cd234778bbe54aecc66-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_JWTService) new() any {
    return &p1579242dd.JWTService{}
}

func (inst* p1579242dd0_impl_JWTService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.JWTService)
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


func (inst*p1579242dd0_impl_JWTService) getRegistryList(ie application.InjectionExt)[]p91f218d46.Registry{
    dst := make([]p91f218d46.Registry, 0)
    src := ie.ListComponents(".class-91f218d46ec21cd234778bbe54aecc66-Registry")
    for _, item1 := range src {
        item2 := item1.(p91f218d46.Registry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p1579242dd0_impl_JWTService) getTokenMaxageDefault(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.default-age-sec}")
}


func (inst*p1579242dd0_impl_JWTService) getTokenMaxageMin(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.min-age-sec}")
}


func (inst*p1579242dd0_impl_JWTService) getTokenMaxageMax(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.max-age-sec}")
}


func (inst*p1579242dd0_impl_JWTService) getSessionMaxageDefault(ie application.InjectionExt)int{
    return ie.GetInt("${security.session.default-age-sec}")
}


func (inst*p1579242dd0_impl_JWTService) getSessionMaxageMin(ie application.InjectionExt)int{
    return ie.GetInt("${security.session.min-age-sec}")
}


func (inst*p1579242dd0_impl_JWTService) getSessionMaxageMax(ie application.InjectionExt)int{
    return ie.GetInt("${security.session.max-age-sec}")
}



// type p1579242dd.SessionServiceImpl in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-SessionServiceImpl
// class:
// alias:alias-d4e0ee677c339b7ffcf1d55767953499-SessionService
// scope:singleton
//
type p1579242dd0_impl_SessionServiceImpl struct {
}

func (inst* p1579242dd0_impl_SessionServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-SessionServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-d4e0ee677c339b7ffcf1d55767953499-SessionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_SessionServiceImpl) new() any {
    return &p1579242dd.SessionServiceImpl{}
}

func (inst* p1579242dd0_impl_SessionServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.SessionServiceImpl)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*p1579242dd0_impl_SessionServiceImpl) getRegistryList(ie application.InjectionExt)[]pd4e0ee677.SessionRegistry{
    dst := make([]pd4e0ee677.SessionRegistry, 0)
    src := ie.ListComponents(".class-d4e0ee677c339b7ffcf1d55767953499-SessionRegistry")
    for _, item1 := range src {
        item2 := item1.(pd4e0ee677.SessionRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p1579242dd.UUIDServiceImpl in package:github.com/starter-go/security/impl
//
// id:com-1579242dd0b3325b-impl-UUIDServiceImpl
// class:
// alias:alias-9621e8b71013b0fc25942a1749ed3652-UUIDService
// scope:singleton
//
type p1579242dd0_impl_UUIDServiceImpl struct {
}

func (inst* p1579242dd0_impl_UUIDServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1579242dd0b3325b-impl-UUIDServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-9621e8b71013b0fc25942a1749ed3652-UUIDService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1579242dd0_impl_UUIDServiceImpl) new() any {
    return &p1579242dd.UUIDServiceImpl{}
}

func (inst* p1579242dd0_impl_UUIDServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1579242dd.UUIDServiceImpl)
	nop(ie, com)

	
    com.Host = inst.getHost(ie)


    return nil
}


func (inst*p1579242dd0_impl_UUIDServiceImpl) getHost(ie application.InjectionExt)string{
    return ie.GetString("${security.uuid.service.hostname}")
}



// type pbf7e6103c.PermissionCacheLoader in package:github.com/starter-go/security/impl/ipermissions
//
// id:com-bf7e6103ccfb71c8-ipermissions-PermissionCacheLoader
// class:
// alias:alias-08935700f7002f152b848e80701dde49-CacheLoader
// scope:singleton
//
type pbf7e6103cc_ipermissions_PermissionCacheLoader struct {
}

func (inst* pbf7e6103cc_ipermissions_PermissionCacheLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bf7e6103ccfb71c8-ipermissions-PermissionCacheLoader"
	r.Classes = ""
	r.Aliases = "alias-08935700f7002f152b848e80701dde49-CacheLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbf7e6103cc_ipermissions_PermissionCacheLoader) new() any {
    return &pbf7e6103c.PermissionCacheLoader{}
}

func (inst* pbf7e6103cc_ipermissions_PermissionCacheLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbf7e6103c.PermissionCacheLoader)
	nop(ie, com)

	
    com.Source = inst.getSource(ie)


    return nil
}


func (inst*pbf7e6103cc_ipermissions_PermissionCacheLoader) getSource(ie application.InjectionExt)[]p08935700f.Registry{
    dst := make([]p08935700f.Registry, 0)
    src := ie.ListComponents(".class-08935700f7002f152b848e80701dde49-Registry")
    for _, item1 := range src {
        item2 := item1.(p08935700f.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type pbf7e6103c.PermissionManagerImpl in package:github.com/starter-go/security/impl/ipermissions
//
// id:com-bf7e6103ccfb71c8-ipermissions-PermissionManagerImpl
// class:
// alias:alias-08935700f7002f152b848e80701dde49-Manager
// scope:singleton
//
type pbf7e6103cc_ipermissions_PermissionManagerImpl struct {
}

func (inst* pbf7e6103cc_ipermissions_PermissionManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bf7e6103ccfb71c8-ipermissions-PermissionManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-08935700f7002f152b848e80701dde49-Manager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbf7e6103cc_ipermissions_PermissionManagerImpl) new() any {
    return &pbf7e6103c.PermissionManagerImpl{}
}

func (inst* pbf7e6103cc_ipermissions_PermissionManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbf7e6103c.PermissionManagerImpl)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*pbf7e6103cc_ipermissions_PermissionManagerImpl) getLoader(ie application.InjectionExt)p08935700f.CacheLoader{
    return ie.GetComponent("#alias-08935700f7002f152b848e80701dde49-CacheLoader").(p08935700f.CacheLoader)
}


