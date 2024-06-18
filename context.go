package security

import "context"

// Context 包含当前上下文中跟安全问题相关的信息
type Context struct {
	Parent context.Context

	// token
	DirtyToken   Token // 表示还没有保存的数据集
	Token        Token
	TokenManager TokenManager

	// session
	DirtySession   Session // 表示还没有保存的数据集
	Session        Session
	SessionManager SessionManager
}

// ContextFactory ...
type ContextFactory interface {
	CreateContext(parent context.Context) (*Context, error)
}

////////////////////////////////////////////////////////////////////////////////

// Flush  保存被标记为 dirty 的数据
func (inst *Context) Flush() error {

	// flush session
	session := inst.DirtySession
	inst.DirtySession = nil
	if session != nil {
		err := inst.SessionManager.Save(inst, session)
		if err != nil {
			return err
		}
	}

	// flush token
	token := inst.DirtyToken
	inst.DirtyToken = nil
	if token != nil {
		err := inst.TokenManager.Save(inst, token)
		if err != nil {
			return err
		}
	}

	return nil
}
