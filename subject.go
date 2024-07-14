package security

// [已废弃]：用 subjects.Subject 代替

// const theSubjectBindingKey = "rbac.Subject#binding"

// // Subject 代表当前主体
// type Subject interface {
// 	GetSession(create bool) (Session, error)

// 	GetToken(create bool) (Token, error)

// 	HasRole(role rbac.RoleName) bool

// 	// 提交已修改的 token & session
// 	Commit() error
// }

// // SetupSubject 把主体绑定到上下文中
// func SetupSubject(c context.Context, ss SessionService) (Subject, error) {

// 	sub, err := GetSubject(c)
// 	if err == nil && sub != nil {
// 		return sub, nil
// 	}

// 	const key = theSubjectBindingKey
// 	kv, err := context2.GetValues(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// make new subject
// 	s2 := &subject{context: c, sessionService: ss}
// 	sub = s2
// 	kv.SetValue(key, sub)
// 	return sub, nil
// }

// // GetSubject 从上下文中提取当前主体
// func GetSubject(c context.Context) (Subject, error) {
// 	const key = theSubjectBindingKey
// 	kv, err := context2.GetValues(c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sub, ok := kv.GetValue(key).(Subject)
// 	if !ok {
// 		return sub, nil
// 	}
// 	return sub, nil
// }

// ////////////////////////////////////////////////////////////////////////////////

// type subject struct {
// 	context        context.Context
// 	session        Session
// 	token          Token
// 	sessionService SessionService
// }

// func (inst *subject) _impl(a Subject) {
// 	a = inst
// }

// func (inst *subject) Init(ss SessionService) {
// 	inst.sessionService = ss
// }

// func (inst *subject) GetSession(create bool) Session {
// 	s := inst.session
// 	if s == nil && create {
// 		ss := inst.sessionService
// 		ctx := inst.context
// 		if ss == nil {
// 			panic("use session without init subject")
// 		}
// 		session, err := ss.GetCurrent(ctx)
// 		if err != nil {
// 			panic(err)
// 		}
// 		s = session
// 		inst.session = session
// 	}
// 	return s
// }

// func (inst *subject) GetToken(create bool) Token {
// }

// func (inst *subject) HasRole(role rbac.RoleName) bool {
// 	str1 := role.String()
// 	str1 = strings.TrimSpace(str1)
// 	if str1 == "" {
// 		return false // 排除空项
// 	}
// 	str1 = strings.ToLower(str1)
// 	session := inst.GetSession(true)
// 	roles := session.Roles().List()
// 	for _, have := range roles {
// 		str2 := have.String()
// 		str2 = strings.TrimSpace(str2)
// 		str2 = strings.ToLower(str2)
// 		if str1 == str2 {
// 			return true
// 		}
// 	}
// 	return false
// }
