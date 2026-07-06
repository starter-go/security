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

    
    inst.register(&p1579242dd0_impl_AuthService1{})
    inst.register(&p1579242dd0_impl_AuthService2{})
    inst.register(&p1579242dd0_impl_AuthenticatorManagerImpl{})
    inst.register(&p1579242dd0_impl_AuthorizerManagerImpl{})
    inst.register(&p1579242dd0_impl_DefaultRandomService{})
    inst.register(&p1579242dd0_impl_DefaultRandomSource{})
    inst.register(&p1579242dd0_impl_JWTCODEC{})
    inst.register(&p1579242dd0_impl_JWTService{})
    inst.register(&p1579242dd0_impl_SessionServiceImpl{})
    inst.register(&p1579242dd0_impl_UUIDServiceImpl{})
    inst.register(&pbf7e6103cc_ipermissions_PermissionCacheLoader{})
    inst.register(&pbf7e6103cc_ipermissions_PermissionManagerImpl{})


    return nil
}
