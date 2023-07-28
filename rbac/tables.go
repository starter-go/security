package rbac

import "github.com/starter-go/libgorm"

const (
	theTableNamespace = "github.com/starter-go/security/rbac"
	theTableGroupName = "rbac"
)

////////////////////////////////////////////////////////////////////////////////

// TableGroup ...
type TableGroup struct {
}

// GetTableGroup ...
func GetTableGroup() *TableGroup {
	t := &TableGroup{}
	return t
}

// Namespace ...
func (inst *TableGroup) Namespace() string {
	return theTableNamespace
}

// Name ...
func (inst *TableGroup) Name() string {
	return theTableGroupName
}

// SetNamer ...
func (inst *TableGroup) SetNamer(namer libgorm.TableNamer) {
	theTableNames.initAll(namer)
}

// Entities ...
func (inst *TableGroup) Entities() []any {
	// todo:list
	list := make([]any, 0)
	list = append(list, &PermissionEntity{})
	list = append(list, &RoleEntity{})
	list = append(list, &UserEntity{})
	list = append(list, &EmailAddressEntity{})
	list = append(list, &PhoneNumberEntity{})
	return list
}

////////////////////////////////////////////////////////////////////////////////

type names struct {
	namer libgorm.TableNamer
	ns    string

	// todo:list
	perm  libgorm.TableNameCache
	role  libgorm.TableNameCache
	user  libgorm.TableNameCache
	email libgorm.TableNameCache
	phone libgorm.TableNameCache
}

var theTableNames names

func (inst *names) initItem(simpleName string, cache *libgorm.TableNameCache) {
	cache.Init(inst.namer, inst.ns, simpleName)
}

func (inst *names) initAll(namer libgorm.TableNamer) {

	inst.ns = theTableNamespace
	inst.namer = namer

	// todo:list
	inst.initItem("permissions", &inst.perm)
	inst.initItem("roles", &inst.role)
	inst.initItem("users", &inst.user)
	inst.initItem("user_email_addresses", &inst.email)
	inst.initItem("user_phone_numbers", &inst.phone)
}

////////////////////////////////////////////////////////////////////////////////
// todo:list

// TableName ...
func (PermissionEntity) TableName() string {
	return theTableNames.perm.Name()
}

// TableName ...
func (RoleEntity) TableName() string {
	return theTableNames.role.Name()
}

// TableName ...
func (UserEntity) TableName() string {
	return theTableNames.user.Name()
}

// TableName ...
func (EmailAddressEntity) TableName() string {
	return theTableNames.email.Name()
}

// TableName ...
func (PhoneNumberEntity) TableName() string {
	return theTableNames.phone.Name()
}
