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

    
    inst.register(&p228d9dd733_users_Convertor{})
    inst.register(&p228d9dd733_users_Dao{})
    inst.register(&p228d9dd733_users_Service{})
    inst.register(&p343202fe35_common_AuthService1{})
    inst.register(&p343202fe35_common_AuthService2{})
    inst.register(&p343202fe35_common_DefaultAuthorizer{})
    inst.register(&p343202fe35_common_JWTCodec{})
    inst.register(&p343202fe35_common_JWTService{})
    inst.register(&p343202fe35_common_PasswordAuth{})
    inst.register(&p343202fe35_common_TableReg{})
    inst.register(&p4f4d46510f_sessions_Service{})
    inst.register(&pa2b13131ae_roles_Convertor{})
    inst.register(&pa2b13131ae_roles_Dao{})
    inst.register(&pa2b13131ae_roles_Service{})
    inst.register(&pb70efa46a6_permissions_Convertor{})
    inst.register(&pb70efa46a6_permissions_Dao{})
    inst.register(&pb70efa46a6_permissions_Service{})


    return nil
}
