package test4security

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

    
    inst.register(&p3197284801_testcom_MockAuth{})
    inst.register(&p3197284801_testcom_MockPermRegDemo{})
    inst.register(&p3197284801_testcom_MockTokenAdapter{})
    inst.register(&p3197284801_testcom_TestCom{})
    inst.register(&p3197284801_testcom_TestRandom{})
    inst.register(&p3197284801_testcom_TestUUID{})


    return nil
}
