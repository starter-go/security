package gen4security
import (
    p3f0a0f8fe "github.com/starter-go/base/lang"
    p512a30914 "github.com/starter-go/libgorm"
    p9d209f7c2 "github.com/starter-go/security/auth"
    p343202fe3 "github.com/starter-go/security/internal/common"
    pb70efa46a "github.com/starter-go/security/internal/permissions"
    pa2b13131a "github.com/starter-go/security/internal/roles"
    p4f4d46510 "github.com/starter-go/security/internal/sessions"
    p228d9dd73 "github.com/starter-go/security/internal/users"
    p91f218d46 "github.com/starter-go/security/jwt"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type p343202fe3.AuthService1 in package:github.com/starter-go/security/internal/common
//
// id:com-343202fe35613f54-common-AuthService1
// class:
// alias:alias-9d209f7c2504d33e6054a2c9998e9485-Service
// scope:singleton
//
type p343202fe35_common_AuthService1 struct {
}

func (inst* p343202fe35_common_AuthService1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-343202fe35613f54-common-AuthService1"
	r.Classes = ""
	r.Aliases = "alias-9d209f7c2504d33e6054a2c9998e9485-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p343202fe35_common_AuthService1) new() any {
    return &p343202fe3.AuthService1{}
}

func (inst* p343202fe35_common_AuthService1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p343202fe3.AuthService1)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*p343202fe35_common_AuthService1) getRegistryList(ie application.InjectionExt)[]p9d209f7c2.Registry{
    dst := make([]p9d209f7c2.Registry, 0)
    src := ie.ListComponents(".class-9d209f7c2504d33e6054a2c9998e9485-Registry")
    for _, item1 := range src {
        item2 := item1.(p9d209f7c2.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type p343202fe3.AuthService2 in package:github.com/starter-go/security/internal/common
//
// id:com-343202fe35613f54-common-AuthService2
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-AuthService
// scope:singleton
//
type p343202fe35_common_AuthService2 struct {
}

func (inst* p343202fe35_common_AuthService2) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-343202fe35613f54-common-AuthService2"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-AuthService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p343202fe35_common_AuthService2) new() any {
    return &p343202fe3.AuthService2{}
}

func (inst* p343202fe35_common_AuthService2) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p343202fe3.AuthService2)
	nop(ie, com)

	
    com.Servic1 = inst.getServic1(ie)


    return nil
}


func (inst*p343202fe35_common_AuthService2) getServic1(ie application.InjectionExt)p9d209f7c2.Service{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-Service").(p9d209f7c2.Service)
}



// type p343202fe3.PasswordAuth in package:github.com/starter-go/security/internal/common
//
// id:com-343202fe35613f54-common-PasswordAuth
// class:class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p343202fe35_common_PasswordAuth struct {
}

func (inst* p343202fe35_common_PasswordAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-343202fe35613f54-common-PasswordAuth"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p343202fe35_common_PasswordAuth) new() any {
    return &p343202fe3.PasswordAuth{}
}

func (inst* p343202fe35_common_PasswordAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p343202fe3.PasswordAuth)
	nop(ie, com)

	
    com.UserDao = inst.getUserDao(ie)


    return nil
}


func (inst*p343202fe35_common_PasswordAuth) getUserDao(ie application.InjectionExt)p2dece1e49.UserDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserDAO").(p2dece1e49.UserDAO)
}



// type p343202fe3.TableReg in package:github.com/starter-go/security/internal/common
//
// id:com-343202fe35613f54-common-TableReg
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry
// alias:
// scope:singleton
//
type p343202fe35_common_TableReg struct {
}

func (inst* p343202fe35_common_TableReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-343202fe35613f54-common-TableReg"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p343202fe35_common_TableReg) new() any {
    return &p343202fe3.TableReg{}
}

func (inst* p343202fe35_common_TableReg) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p343202fe3.TableReg)
	nop(ie, com)

	


    return nil
}



// type p343202fe3.DefaultAuthorizer in package:github.com/starter-go/security/internal/common
//
// id:com-343202fe35613f54-common-DefaultAuthorizer
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p343202fe35_common_DefaultAuthorizer struct {
}

func (inst* p343202fe35_common_DefaultAuthorizer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-343202fe35613f54-common-DefaultAuthorizer"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p343202fe35_common_DefaultAuthorizer) new() any {
    return &p343202fe3.DefaultAuthorizer{}
}

func (inst* p343202fe35_common_DefaultAuthorizer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p343202fe3.DefaultAuthorizer)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.MaxAge = inst.getMaxAge(ie)


    return nil
}


func (inst*p343202fe35_common_DefaultAuthorizer) getService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p343202fe35_common_DefaultAuthorizer) getMaxAge(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.max-age}")
}



// type p343202fe3.JWTCodec in package:github.com/starter-go/security/internal/common
//
// id:com-343202fe35613f54-common-JWTCodec
// class:class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type p343202fe35_common_JWTCodec struct {
}

func (inst* p343202fe35_common_JWTCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-343202fe35613f54-common-JWTCodec"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p343202fe35_common_JWTCodec) new() any {
    return &p343202fe3.JWTCodec{}
}

func (inst* p343202fe35_common_JWTCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p343202fe3.JWTCodec)
	nop(ie, com)

	
    com.Secret = inst.getSecret(ie)


    return nil
}


func (inst*p343202fe35_common_JWTCodec) getSecret(ie application.InjectionExt)string{
    return ie.GetString("${security.jwt.secret}")
}



// type p343202fe3.JWTService in package:github.com/starter-go/security/internal/common
//
// id:com-343202fe35613f54-common-JWTService
// class:
// alias:alias-91f218d46ec21cd234778bbe54aecc66-Service
// scope:singleton
//
type p343202fe35_common_JWTService struct {
}

func (inst* p343202fe35_common_JWTService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-343202fe35613f54-common-JWTService"
	r.Classes = ""
	r.Aliases = "alias-91f218d46ec21cd234778bbe54aecc66-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p343202fe35_common_JWTService) new() any {
    return &p343202fe3.JWTService{}
}

func (inst* p343202fe35_common_JWTService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p343202fe3.JWTService)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*p343202fe35_common_JWTService) getRegistryList(ie application.InjectionExt)[]p91f218d46.Registry{
    dst := make([]p91f218d46.Registry, 0)
    src := ie.ListComponents(".class-91f218d46ec21cd234778bbe54aecc66-Registry")
    for _, item1 := range src {
        item2 := item1.(p91f218d46.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type pb70efa46a.Convertor in package:github.com/starter-go/security/internal/permissions
//
// id:com-b70efa46a69690e8-permissions-Convertor
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionConvertor
// scope:singleton
//
type pb70efa46a6_permissions_Convertor struct {
}

func (inst* pb70efa46a6_permissions_Convertor) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b70efa46a69690e8-permissions-Convertor"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb70efa46a6_permissions_Convertor) new() any {
    return &pb70efa46a.Convertor{}
}

func (inst* pb70efa46a6_permissions_Convertor) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb70efa46a.Convertor)
	nop(ie, com)

	


    return nil
}



// type pb70efa46a.Dao in package:github.com/starter-go/security/internal/permissions
//
// id:com-b70efa46a69690e8-permissions-Dao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionDAO
// scope:singleton
//
type pb70efa46a6_permissions_Dao struct {
}

func (inst* pb70efa46a6_permissions_Dao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b70efa46a69690e8-permissions-Dao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb70efa46a6_permissions_Dao) new() any {
    return &pb70efa46a.Dao{}
}

func (inst* pb70efa46a6_permissions_Dao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb70efa46a.Dao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenerator = inst.getUUIDGenerator(ie)


    return nil
}


func (inst*pb70efa46a6_permissions_Dao) getAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*pb70efa46a6_permissions_Dao) getUUIDGenerator(ie application.InjectionExt)p3f0a0f8fe.UUIDGenerator{
    return ie.GetComponent("#alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator").(p3f0a0f8fe.UUIDGenerator)
}



// type pb70efa46a.Service in package:github.com/starter-go/security/internal/permissions
//
// id:com-b70efa46a69690e8-permissions-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionService
// scope:singleton
//
type pb70efa46a6_permissions_Service struct {
}

func (inst* pb70efa46a6_permissions_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b70efa46a69690e8-permissions-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb70efa46a6_permissions_Service) new() any {
    return &pb70efa46a.Service{}
}

func (inst* pb70efa46a6_permissions_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb70efa46a.Service)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Convertor = inst.getConvertor(ie)


    return nil
}


func (inst*pb70efa46a6_permissions_Service) getDao(ie application.InjectionExt)p2dece1e49.PermissionDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionDAO").(p2dece1e49.PermissionDAO)
}


func (inst*pb70efa46a6_permissions_Service) getConvertor(ie application.InjectionExt)p2dece1e49.PermissionConvertor{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionConvertor").(p2dece1e49.PermissionConvertor)
}



// type pa2b13131a.Convertor in package:github.com/starter-go/security/internal/roles
//
// id:com-a2b13131aed7b874-roles-Convertor
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-RoleConvertor
// scope:singleton
//
type pa2b13131ae_roles_Convertor struct {
}

func (inst* pa2b13131ae_roles_Convertor) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a2b13131aed7b874-roles-Convertor"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-RoleConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa2b13131ae_roles_Convertor) new() any {
    return &pa2b13131a.Convertor{}
}

func (inst* pa2b13131ae_roles_Convertor) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa2b13131a.Convertor)
	nop(ie, com)

	


    return nil
}



// type pa2b13131a.Dao in package:github.com/starter-go/security/internal/roles
//
// id:com-a2b13131aed7b874-roles-Dao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-RoleDAO
// scope:singleton
//
type pa2b13131ae_roles_Dao struct {
}

func (inst* pa2b13131ae_roles_Dao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a2b13131aed7b874-roles-Dao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-RoleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa2b13131ae_roles_Dao) new() any {
    return &pa2b13131a.Dao{}
}

func (inst* pa2b13131ae_roles_Dao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa2b13131a.Dao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenerator = inst.getUUIDGenerator(ie)


    return nil
}


func (inst*pa2b13131ae_roles_Dao) getAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*pa2b13131ae_roles_Dao) getUUIDGenerator(ie application.InjectionExt)p3f0a0f8fe.UUIDGenerator{
    return ie.GetComponent("#alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator").(p3f0a0f8fe.UUIDGenerator)
}



// type pa2b13131a.Service in package:github.com/starter-go/security/internal/roles
//
// id:com-a2b13131aed7b874-roles-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-RoleService
// scope:singleton
//
type pa2b13131ae_roles_Service struct {
}

func (inst* pa2b13131ae_roles_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a2b13131aed7b874-roles-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-RoleService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa2b13131ae_roles_Service) new() any {
    return &pa2b13131a.Service{}
}

func (inst* pa2b13131ae_roles_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa2b13131a.Service)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Convertor = inst.getConvertor(ie)


    return nil
}


func (inst*pa2b13131ae_roles_Service) getDao(ie application.InjectionExt)p2dece1e49.RoleDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleDAO").(p2dece1e49.RoleDAO)
}


func (inst*pa2b13131ae_roles_Service) getConvertor(ie application.InjectionExt)p2dece1e49.RoleConvertor{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleConvertor").(p2dece1e49.RoleConvertor)
}



// type p4f4d46510.Service in package:github.com/starter-go/security/internal/sessions
//
// id:com-4f4d46510fdfb358-sessions-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-SessionService
// scope:singleton
//
type p4f4d46510f_sessions_Service struct {
}

func (inst* p4f4d46510f_sessions_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4f4d46510fdfb358-sessions-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-SessionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4f4d46510f_sessions_Service) new() any {
    return &p4f4d46510.Service{}
}

func (inst* p4f4d46510f_sessions_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4f4d46510.Service)
	nop(ie, com)

	
    com.JWTser = inst.getJWTser(ie)


    return nil
}


func (inst*p4f4d46510f_sessions_Service) getJWTser(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type p228d9dd73.Convertor in package:github.com/starter-go/security/internal/users
//
// id:com-228d9dd73338a2c5-users-Convertor
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-UserConvertor
// scope:singleton
//
type p228d9dd733_users_Convertor struct {
}

func (inst* p228d9dd733_users_Convertor) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-228d9dd73338a2c5-users-Convertor"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-UserConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p228d9dd733_users_Convertor) new() any {
    return &p228d9dd73.Convertor{}
}

func (inst* p228d9dd733_users_Convertor) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p228d9dd73.Convertor)
	nop(ie, com)

	


    return nil
}



// type p228d9dd73.Dao in package:github.com/starter-go/security/internal/users
//
// id:com-228d9dd73338a2c5-users-Dao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-UserDAO
// scope:singleton
//
type p228d9dd733_users_Dao struct {
}

func (inst* p228d9dd733_users_Dao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-228d9dd73338a2c5-users-Dao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-UserDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p228d9dd733_users_Dao) new() any {
    return &p228d9dd73.Dao{}
}

func (inst* p228d9dd733_users_Dao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p228d9dd73.Dao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenerator = inst.getUUIDGenerator(ie)


    return nil
}


func (inst*p228d9dd733_users_Dao) getAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*p228d9dd733_users_Dao) getUUIDGenerator(ie application.InjectionExt)p3f0a0f8fe.UUIDGenerator{
    return ie.GetComponent("#alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator").(p3f0a0f8fe.UUIDGenerator)
}



// type p228d9dd73.Service in package:github.com/starter-go/security/internal/users
//
// id:com-228d9dd73338a2c5-users-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-UserService
// scope:singleton
//
type p228d9dd733_users_Service struct {
}

func (inst* p228d9dd733_users_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-228d9dd73338a2c5-users-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-UserService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p228d9dd733_users_Service) new() any {
    return &p228d9dd73.Service{}
}

func (inst* p228d9dd733_users_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p228d9dd73.Service)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Convertor = inst.getConvertor(ie)


    return nil
}


func (inst*p228d9dd733_users_Service) getDao(ie application.InjectionExt)p2dece1e49.UserDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserDAO").(p2dece1e49.UserDAO)
}


func (inst*p228d9dd733_users_Service) getConvertor(ie application.InjectionExt)p2dece1e49.UserConvertor{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserConvertor").(p2dece1e49.UserConvertor)
}


