package rbac

import (
	"context"
	"strconv"
)

// UserID 是通用的用户标识符
type UserID int64

// UserName 表示用户名
type UserName string

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
	Enabled  bool         `json:"enabled"`
}

// UserQuery 是 User 的查询参数
type UserQuery struct {
	Pagination Pagination
	Conditions map[string]string
	All        bool // 查询全部条目
}

// UserService 是针对 UserDTO 的服务
type UserService interface {
	Insert(c context.Context, o *UserDTO) (*UserDTO, error)
	Update(c context.Context, id UserID, o *UserDTO) (*UserDTO, error)
	Delete(c context.Context, id UserID) error

	Find(c context.Context, id UserID) (*UserDTO, error)
	List(c context.Context, q *UserQuery) ([]*UserDTO, error)
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

func (name UserName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////
