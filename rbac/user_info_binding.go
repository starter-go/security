package rbac

import (
	"context"
	"fmt"
)

const theUserInfoBindingKey = "rbac.UserInfoBinding#binding"

// UserInfo 包含当前用户的基本信息
type UserInfo struct {
	UserDTO
}

// UserInfoBinding 负责把用户信息绑定到当前上下文
type UserInfoBinding struct {
	key  string
	info UserInfo
}

// Key 取绑定键名
func (inst *UserInfoBinding) Key() string {
	return inst.key
}

// UserInfo 取用户信息
func (inst *UserInfoBinding) UserInfo() *UserInfo {
	return &inst.info
}

// GetUserInfo 从上下文中获取已绑定的用户信息
func GetUserInfo(cc context.Context) (*UserInfo, error) {
	const key = theUserInfoBindingKey
	o := cc.Value(key)
	binding, ok := o.(*UserInfoBinding)
	if !ok {
		return nil, fmt.Errorf("no rbac.UserInfo bound with the context")
	}
	return binding.UserInfo(), nil
}

// NewUserInfoBinding 新建并初始化一个 UserInfoBinding
func NewUserInfoBinding(info *UserInfo) *UserInfoBinding {
	b := &UserInfoBinding{}
	b.key = theUserInfoBindingKey
	if info != nil {
		b.info = *info
	}
	return b
}
