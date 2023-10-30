package rbac

// Conditions 表示查询条件, 取值参考 gorm.DB.Where()
type Conditions struct {
	Query string   `json:"query"`
	Args  []string `json:"args"`
}
