package internal

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/jwt"
)

// RbacTokenServiceImpl ... 通过 jwt 实现 rbac.TokenService 的功能
type RbacTokenServiceImpl struct {

	//starter:component

	_as func(rbac.TokenService) //starter:as("#")

	JWTService jwt.Service //starter:inject("#")

}

func (inst *RbacTokenServiceImpl) _impl() rbac.TokenService { return inst }

// GetCurrent ...
func (inst *RbacTokenServiceImpl) GetCurrent(c context.Context) (*rbac.TokenDTO, error) {

	ser := inst.JWTService
	src, err := ser.GetDTO(c)
	if err != nil {
		return nil, err
	}

	dst := new(rbac.TokenDTO)
	dst.CurrentUser = src.CurrentUser
	dst.Session = src.Session
	return dst, nil
}

// PutCurrent ...
func (inst *RbacTokenServiceImpl) PutCurrent(c context.Context, token *rbac.TokenDTO) (*rbac.TokenDTO, error) {

	ser := inst.JWTService
	src := token
	dst := new(jwt.Token)

	dst.CurrentUser = src.CurrentUser
	dst.Session = src.Session

	err := ser.SetDTO(c, dst)
	if err != nil {
		return nil, err
	}
	return src, nil
}
