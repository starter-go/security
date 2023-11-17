package gen4security
import (
    pd4e0ee677 "github.com/starter-go/security"
    p9d209f7c2 "github.com/starter-go/security/auth"
    pf41d62225 "github.com/starter-go/security/internal"
    p3aad4dd5c "github.com/starter-go/security/internal/crypts"
    p3a8e2a16f "github.com/starter-go/security/internal/crypts/aes"
    paa12a345b "github.com/starter-go/security/internal/crypts/rsa"
    p91f218d46 "github.com/starter-go/security/jwt"
    p5d8e1a661 "github.com/starter-go/security/keys"
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


    return nil
}


func (inst*pf41d62225c_internal_AuthService1) getAuthenticators(ie application.InjectionExt)p9d209f7c2.AuthenticatorManager{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-AuthenticatorManager").(p9d209f7c2.AuthenticatorManager)
}


func (inst*pf41d62225c_internal_AuthService1) getAuthorizers(ie application.InjectionExt)p9d209f7c2.AuthorizerManager{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-AuthorizerManager").(p9d209f7c2.AuthorizerManager)
}



// type pf41d62225.AuthService2 in package:github.com/starter-go/security/internal
//
// id:com-f41d62225c42aa4c-internal-AuthService2
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-AuthService
// scope:singleton
//
type pf41d62225c_internal_AuthService2 struct {
}

func (inst* pf41d62225c_internal_AuthService2) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f41d62225c42aa4c-internal-AuthService2"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-AuthService"
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



// type p3aad4dd5c.ManagerImpl in package:github.com/starter-go/security/internal/crypts
//
// id:com-3aad4dd5c4bff300-crypts-ManagerImpl
// class:
// alias:alias-5d8e1a661f387d56d217edd5cab8802a-Manager
// scope:singleton
//
type p3aad4dd5c4_crypts_ManagerImpl struct {
}

func (inst* p3aad4dd5c4_crypts_ManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3aad4dd5c4bff300-crypts-ManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-5d8e1a661f387d56d217edd5cab8802a-Manager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3aad4dd5c4_crypts_ManagerImpl) new() any {
    return &p3aad4dd5c.ManagerImpl{}
}

func (inst* p3aad4dd5c4_crypts_ManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3aad4dd5c.ManagerImpl)
	nop(ie, com)

	
    com.Regs = inst.getRegs(ie)


    return nil
}


func (inst*p3aad4dd5c4_crypts_ManagerImpl) getRegs(ie application.InjectionExt)[]p5d8e1a661.Registry{
    dst := make([]p5d8e1a661.Registry, 0)
    src := ie.ListComponents(".class-5d8e1a661f387d56d217edd5cab8802a-Registry")
    for _, item1 := range src {
        item2 := item1.(p5d8e1a661.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type p3aad4dd5c.ServiceImpl in package:github.com/starter-go/security/internal/crypts
//
// id:com-3aad4dd5c4bff300-crypts-ServiceImpl
// class:
// alias:alias-5d8e1a661f387d56d217edd5cab8802a-Service
// scope:singleton
//
type p3aad4dd5c4_crypts_ServiceImpl struct {
}

func (inst* p3aad4dd5c4_crypts_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3aad4dd5c4bff300-crypts-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-5d8e1a661f387d56d217edd5cab8802a-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3aad4dd5c4_crypts_ServiceImpl) new() any {
    return &p3aad4dd5c.ServiceImpl{}
}

func (inst* p3aad4dd5c4_crypts_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3aad4dd5c.ServiceImpl)
	nop(ie, com)

	
    com.Manager = inst.getManager(ie)


    return nil
}


func (inst*p3aad4dd5c4_crypts_ServiceImpl) getManager(ie application.InjectionExt)p5d8e1a661.Manager{
    return ie.GetComponent("#alias-5d8e1a661f387d56d217edd5cab8802a-Manager").(p5d8e1a661.Manager)
}



// type p3a8e2a16f.Provider in package:github.com/starter-go/security/internal/crypts/aes
//
// id:com-3a8e2a16fe53bb90-aes-Provider
// class:class-5d8e1a661f387d56d217edd5cab8802a-Registry
// alias:
// scope:singleton
//
type p3a8e2a16fe_aes_Provider struct {
}

func (inst* p3a8e2a16fe_aes_Provider) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a8e2a16fe53bb90-aes-Provider"
	r.Classes = "class-5d8e1a661f387d56d217edd5cab8802a-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a8e2a16fe_aes_Provider) new() any {
    return &p3a8e2a16f.Provider{}
}

func (inst* p3a8e2a16fe_aes_Provider) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a8e2a16f.Provider)
	nop(ie, com)

	
    com.Rand = inst.getRand(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Priority = inst.getPriority(ie)


    return nil
}


func (inst*p3a8e2a16fe_aes_Provider) getRand(ie application.InjectionExt)p9621e8b71.Service{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-Service").(p9621e8b71.Service)
}


func (inst*p3a8e2a16fe_aes_Provider) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${security.algorithms.aes.enabled}")
}


func (inst*p3a8e2a16fe_aes_Provider) getPriority(ie application.InjectionExt)int{
    return ie.GetInt("${security.algorithms.aes.priority}")
}



// type paa12a345b.Provider in package:github.com/starter-go/security/internal/crypts/rsa
//
// id:com-aa12a345bda7a8c2-rsa-Provider
// class:class-5d8e1a661f387d56d217edd5cab8802a-Registry
// alias:
// scope:singleton
//
type paa12a345bd_rsa_Provider struct {
}

func (inst* paa12a345bd_rsa_Provider) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-aa12a345bda7a8c2-rsa-Provider"
	r.Classes = "class-5d8e1a661f387d56d217edd5cab8802a-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* paa12a345bd_rsa_Provider) new() any {
    return &paa12a345b.Provider{}
}

func (inst* paa12a345bd_rsa_Provider) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*paa12a345b.Provider)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Priority = inst.getPriority(ie)


    return nil
}


func (inst*paa12a345bd_rsa_Provider) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${security.algorithms.rsa.enabled}")
}


func (inst*paa12a345bd_rsa_Provider) getPriority(ie application.InjectionExt)int{
    return ie.GetInt("${security.algorithms.rsa.priority}")
}



// type paa12a345b.SignWithRsaProvider in package:github.com/starter-go/security/internal/crypts/rsa
//
// id:com-aa12a345bda7a8c2-rsa-SignWithRsaProvider
// class:class-5d8e1a661f387d56d217edd5cab8802a-Registry
// alias:
// scope:singleton
//
type paa12a345bd_rsa_SignWithRsaProvider struct {
}

func (inst* paa12a345bd_rsa_SignWithRsaProvider) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-aa12a345bda7a8c2-rsa-SignWithRsaProvider"
	r.Classes = "class-5d8e1a661f387d56d217edd5cab8802a-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* paa12a345bd_rsa_SignWithRsaProvider) new() any {
    return &paa12a345b.SignWithRsaProvider{}
}

func (inst* paa12a345bd_rsa_SignWithRsaProvider) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*paa12a345b.SignWithRsaProvider)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Priority = inst.getPriority(ie)


    return nil
}


func (inst*paa12a345bd_rsa_SignWithRsaProvider) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${security.algorithms.rsa.enabled}")
}


func (inst*paa12a345bd_rsa_SignWithRsaProvider) getPriority(ie application.InjectionExt)int{
    return ie.GetInt("${security.algorithms.rsa.priority}")
}


