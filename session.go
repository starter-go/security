package security

import (
	"context"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac"
)

// Session 代表当前会话
type Session interface {
	GetProperties() properties.Table

	Get() *rbac.SessionDTO

	Set(s *rbac.SessionDTO)

	UserID() rbac.UserID

	UserName() rbac.UserName

	Nickname() string

	Avatar() string

	Roles() rbac.RoleNameList

	Authenticated() bool
}

// SessionService 是针对 SessionDTO 的服务
type SessionService interface {
	GetCurrent(c context.Context) (Session, error)
}

// SessionProvider  会话的实现方案
type SessionProvider interface {
	Support(c context.Context) bool
	Current(c context.Context) (Session, error)
}

// SessionRegistration 会话方案的注册信息
type SessionRegistration struct {
	Name     string
	Enabled  bool
	Priority int
	Provider SessionProvider
}

// SessionRegistry 会话方案的注册接口
type SessionRegistry interface {
	Registration() *SessionRegistration
}

// SessionGetter ...
type SessionGetter interface {
	Get(ctx *Context) (Session, error)
}

// SessionLoader ...
type SessionLoader interface {
	Load(ctx *Context) (Session, error)
}

// SessionSaver ...
type SessionSaver interface {
	Save(ctx *Context, se Session) error
}

// SessionSetter ...
type SessionSetter interface {
	Set(ctx *Context, se Session) error
}

// SessionManager ...
type SessionManager interface {
	SessionGetter
	SessionSetter
	SessionLoader
	SessionSaver
}
