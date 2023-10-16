package random

import "github.com/starter-go/base/lang"

////////////////////////////////////////////////////////////////////////////////

// UUIDBuilder ...
type UUIDBuilder interface {

	// Nonce(value string)
	// Host(value string)
	// StartedAt(value lang.Time) // 服务启动时间戳
	// CreatedAt(value lang.Time) // UUID 创建时间戳
	// SN(value int64)            // UUID 序列号
	// Service(value UUIDService)

	Generate() lang.UUID

	Others(value ...string) UUIDBuilder // 其它参数
	Class(value string) UUIDBuilder
	ID(value string) UUIDBuilder
}

////////////////////////////////////////////////////////////////////////////////

// UUIDService 是 UUID 创建服务
type UUIDService interface {
	Build() UUIDBuilder
}

////////////////////////////////////////////////////////////////////////////////
