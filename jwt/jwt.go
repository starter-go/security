package jwt

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// Text 是文本形式的 JWT
type Text string

// DTO 是 JSON 结构形式的 JWT
type DTO struct {
	rbac.UserDTO
	Properties map[string]string `json:"properties"`
}

// Getter 用来获取跟上下文绑定的JWT
type Getter interface {
	GetDTO(c context.Context) (*DTO, error)
	GetText(c context.Context) (Text, error)
}

// Setter 用来设置跟上下文绑定的JWT
type Setter interface {
	SetDTO(c context.Context, o *DTO) error
	SetText(c context.Context, t *Text) error
}

// CODEC 是 JWT 的编解码器
type CODEC interface {
	Encode(o *DTO) (Text, error)
	Decode(t Text) (*DTO, error)
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
}
