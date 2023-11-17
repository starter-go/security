package gen4security

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p3a8e2a16fe_aes_Provider{})
    inst.register(&p3aad4dd5c4_crypts_ManagerImpl{})
    inst.register(&p3aad4dd5c4_crypts_ServiceImpl{})
    inst.register(&paa12a345bd_rsa_Provider{})
    inst.register(&paa12a345bd_rsa_SignWithRsaProvider{})
    inst.register(&pf41d62225c_internal_AuthService1{})
    inst.register(&pf41d62225c_internal_AuthService2{})
    inst.register(&pf41d62225c_internal_AuthenticatorManagerImpl{})
    inst.register(&pf41d62225c_internal_AuthorizerManagerImpl{})
    inst.register(&pf41d62225c_internal_DefaultRandomService{})
    inst.register(&pf41d62225c_internal_DefaultRandomSource{})
    inst.register(&pf41d62225c_internal_JWTCODEC{})
    inst.register(&pf41d62225c_internal_JWTService{})
    inst.register(&pf41d62225c_internal_SessionServiceImpl{})
    inst.register(&pf41d62225c_internal_UUIDServiceImpl{})


    return nil
}
