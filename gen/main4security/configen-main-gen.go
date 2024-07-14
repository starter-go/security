package main4security

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

    
    inst.register(&pf41d62225c_internal_AuthService1{})
    inst.register(&pf41d62225c_internal_AuthService2{})
    inst.register(&pf41d62225c_internal_AuthenticatorManagerImpl{})
    inst.register(&pf41d62225c_internal_AuthorizerManagerImpl{})
    inst.register(&pf41d62225c_internal_DefaultRandomService{})
    inst.register(&pf41d62225c_internal_DefaultRandomSource{})
    inst.register(&pf41d62225c_internal_DefaultSubjectLoader{})
    inst.register(&pf41d62225c_internal_JWTCODEC{})
    inst.register(&pf41d62225c_internal_JWTService{})
    inst.register(&pf41d62225c_internal_RbacSessionServiceImpl{})
    inst.register(&pf41d62225c_internal_RbacSubjectServiceImpl{})
    inst.register(&pf41d62225c_internal_RbacTokenServiceImpl{})
    inst.register(&pf41d62225c_internal_SessionServiceImpl{})
    inst.register(&pf41d62225c_internal_UUIDServiceImpl{})


    return nil
}
