package jwt

import "github.com/starter-go/security/rbac"

// Text ...
type Text string

// DTO ...
type DTO struct {
	rbac.UserDTO

	Properties map[string]string
}

// Service 是 JWT 的编解码服务
type Service interface {
	Encode(o *DTO) (Text, error)
	Decode(t Text) (*DTO, error)
}
