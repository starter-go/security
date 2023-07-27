package rbac

// Pagination 是通用的分页参数
type Pagination struct {
	Page  int64 `json:"page"`  // 页码, first=1
	Size  int   `json:"size"`  // 页大小
	Total int64 `json:"total"` // 所有页面的条目总数
}

// Limit 用于 SQL 查询的页面大小
func (inst *Pagination) Limit() int64 {
	const min = 1
	limit := inst.Size
	if limit < min {
		limit = min
	}
	return int64(limit)
}

// Offset 用于 SQL 查询的条目位置
func (inst *Pagination) Offset() int64 {
	const min = 0
	i := inst.Page - 1
	if i < min {
		i = min
	}
	size := inst.Limit()
	return (i * size)
}
