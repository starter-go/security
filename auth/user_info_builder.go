package auth

// // UserBuilder 用于创建 User 实例
// type UserBuilder struct {
// 	ID          rbac.UserID
// 	Name        rbac.UserName
// 	Roles       rbac.RoleNameList
// 	Avatar      string
// 	DisplayName string
// }

// // Create 。。。
// func (inst *UserBuilder) Create() User {
// 	info := &userInfo{
// 		avatar:   inst.Avatar,
// 		id:       inst.ID,
// 		name:     inst.Name,
// 		nickname: inst.DisplayName,
// 		roles:    inst.Roles,
// 	}
// 	return info
// }

// ////////////////////////////////////////////////////////////////////////////////

// type userInfo struct {
// 	id       rbac.UserID
// 	name     rbac.UserName
// 	roles    rbac.RoleNameList
// 	avatar   string
// 	nickname string
// }

// func (inst *userInfo) _impl() User {
// 	return inst
// }

// func (inst *userInfo) Name() rbac.UserName {
// 	return inst.name
// }

// func (inst *userInfo) ID() rbac.UserID {
// 	return inst.id
// }

// func (inst *userInfo) Avatar() string {
// 	return inst.avatar
// }

// func (inst *userInfo) DisplayName() string {
// 	return inst.nickname
// }

// func (inst *userInfo) Roles() rbac.RoleNameList {
// 	return inst.roles
// }
