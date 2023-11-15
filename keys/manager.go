package keys

// Manager 是密钥算法管理器
type Manager interface {

	// 根据名称，查找算法（可能有多条结果）
	Find(algorithm string) ([]Algorithm, error)

	// 根据名称，查找算法（如果选择器函数为 nil, 则适配任何条目）
	Get(algorithm string, selector func(reg *Registration) bool) (Algorithm, error)
}
