package gen4test

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

    
    inst.register(&p5772338936_code_MockAuth{})
    inst.register(&p5772338936_code_MockTokenAdapter{})
    inst.register(&p5772338936_code_TestAES{})
    inst.register(&p5772338936_code_TestCom{})
    inst.register(&p5772338936_code_TestRSA{})
    inst.register(&p5772338936_code_TestRandom{})
    inst.register(&p5772338936_code_TestUUID{})


    return nil
}
