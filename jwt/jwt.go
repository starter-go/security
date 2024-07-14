package jwt

import (
	"context"

	"github.com/starter-go/rbac"
)

// Text 是文本形式的 JWT
type Text string

// Token 是 JSON 结构形式的 JWT
type Token struct {

	// 当前用户信息
	rbac.CurrentUser

	Session rbac.SessionID `json:"session"` // 会话的 UUID
}

// Getter 用来获取跟上下文绑定的JWT
type Getter interface {
	GetText(c context.Context) (Text, error)
}

// Setter 用来设置跟上下文绑定的JWT
type Setter interface {
	SetText(c context.Context, t Text) error
}

// Encoder 是 JWT 的编码器
type Encoder interface {
	Encode(o *Token) (Text, error)
}

// Decoder 是 JWT 的解码器
type Decoder interface {
	Decode(t Text) (*Token, error)
}

// CODEC 是 JWT 的编解码器
type CODEC interface {
	Encoder
	Decoder
}

// Adapter 用来适配各种不同的上下文
type Adapter interface {
	Getter
	Setter
	Accept(c context.Context) bool
}

// Registration 表示 JWT 组件的注册信息
type Registration struct {
	Adapter  Adapter
	CODEC    CODEC
	Priority int
	Enabled  bool
}

// Registry 是 JWT 组件的注册接口
type Registry interface {
	ListRegistrations() []*Registration
}

// Service 提供全部功能的 JWT 服务
type Service interface {
	Getter
	Setter
	CODEC

	GetDTO(c context.Context) (*Token, error)
	SetDTO(c context.Context, o *Token) error
}

////////////////////////////////////////////////////////////////////////////////

func (t Text) String() string {
	return string(t)
}
