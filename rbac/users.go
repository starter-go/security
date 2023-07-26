package rbac

import (
	"context"
	"strconv"

	"github.com/starter-go/base/lang"
	"gorm.io/gorm"
)

// UserID 是通用的用户标识符
type UserID int64

// UserName 表示用户名
type UserName string

// UserEntity 表示 User 的数据库实体
type UserEntity struct {
	ID UserID

	BaseEntity

	Name     UserName     // 用户名
	Nickname string       // 昵称
	Avatar   string       // 头像 (HTTP-URL)
	Phone    string       // 主要的手机号
	Email    string       // 主要的 e-mail 地址
	Roles    RoleNameList // 用户的角色

	Password lang.Hex // 用户当前的密码
	Salt     lang.Hex // 跟密码相关的盐
}

// UserDTO 表示 User 的 REST 网络对象
type UserDTO struct {
	ID UserID `json:"id"`

	BaseDTO

	Name     UserName     `json:"name"`
	NickName string       `json:"nickname"`
	Avatar   string       `json:"avatar"`
	Phone    string       `json:"phone"`
	Email    string       `json:"email"`
	Roles    RoleNameList `json:"roles"`
}

// UserQuery 是 User 的查询参数
type UserQuery struct {
	Pagination Pagination
}

// UserService 是针对 UserDTO 的服务
type UserService interface {
	Insert(c context.Context, o *UserDTO) (*UserDTO, error)
	Update(c context.Context, id UserID, o *UserDTO) (*UserDTO, error)
	Delete(c context.Context, id UserID) error

	Find(c context.Context, id UserID) (*UserDTO, error)
	List(c context.Context, q *UserQuery) ([]*UserDTO, error)
}

// UserDAO 是数据库访问对象
type UserDAO interface {
	Insert(db *gorm.DB, o *UserEntity) (*UserEntity, error)
	Update(db *gorm.DB, id UserID, updater func(o *UserEntity)) (*UserEntity, error)
	Delete(db *gorm.DB, id UserID) error

	FindByAny(db *gorm.DB, text string) (*UserEntity, error)
	FindByName(db *gorm.DB, name string) (*UserEntity, error)
	FindByPhone(db *gorm.DB, phone string) (*UserEntity, error)
	FindByEmail(db *gorm.DB, email string) (*UserEntity, error)

	Find(db *gorm.DB, id UserID) (*UserEntity, error)
	List(db *gorm.DB, q *UserQuery) ([]*UserEntity, error)
}

// UserConvertor 负责 dto <==> entity 的转换
type UserConvertor interface {
	ConvertE2D(c context.Context, entity *UserEntity) (*UserDTO, error)
	ConvertD2E(c context.Context, dto *UserDTO) (*UserEntity, error)

	ConvertListE2D(c context.Context, entity []*UserEntity) ([]*UserDTO, error)
}

////////////////////////////////////////////////////////////////////////////////

func (id UserID) String() string {
	n := int64(id)
	return strconv.FormatInt(n, 10)
}

// ParseUserID 把字符串解析为 UserID
func ParseUserID(s string) (UserID, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return UserID(n), nil
}

////////////////////////////////////////////////////////////////////////////////
