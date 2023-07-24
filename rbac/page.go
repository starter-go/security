package rbac

// Pagination 是通用的分页参数
type Pagination struct {
	Number int64 `json:"number"` // 页码, first=1
	Size   int   `json:"size"`   // 页大小
	Total  int64 `json:"total"`  // 所有页面的条目总数
}

// Limit 用于 SQL 查询的页面大小
func (inst *Pagination) Limit() int {
	const min = 1
	limit := inst.Size
	if limit < min {
		limit = min
	}
	return limit
}

// Offset 用于 SQL 查询的条目位置
func (inst *Pagination) Offset() int64 {
	i := inst.Number - 1
	size := int64(inst.Size)
	return (i * size)
}
