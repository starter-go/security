package crypts

import (
	"fmt"
	"sort"
	"strings"

	"github.com/starter-go/security/keys"
)

// ManagerImpl  实现 keys.Manager
type ManagerImpl struct {

	//starter:component
	_as func(keys.Manager) //starter:as("#")

	Regs []keys.Registry //starter:inject(".")

	items []*keys.Registration // the cached items
}

func (inst *ManagerImpl) _impl() keys.Manager {
	return inst
}

func (inst *ManagerImpl) loadAll() []*keys.Registration {
	src := inst.Regs
	dst := make([]*keys.Registration, 0)
	for _, r1 := range src {
		r2list := r1.ListRegistrations()
		for _, r2 := range r2list {
			alg := r2.Algorithm
			if !r2.Enabled || alg == nil {
				continue
			}
			name := alg.Name()
			r2.Name = inst.normalizeName(name)
			r2.Provider = alg.Provider()
			dst = append(dst, r2)
		}
	}
	return inst.sort(dst)
}

func (inst *ManagerImpl) getAll() []*keys.Registration {
	c := inst.items
	if c == nil {
		c = inst.loadAll()
		inst.items = c
	}
	return c
}

func (inst *ManagerImpl) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}

// Find 根据名称，查找算法（可能有多条结果）
func (inst *ManagerImpl) Find(algorithm string) ([]keys.Algorithm, error) {
	dst := make([]keys.Algorithm, 0)
	algorithm = inst.normalizeName(algorithm)
	all := inst.getAll()
	for _, item := range all {
		if item.Name == algorithm {
			dst = append(dst, item.Algorithm)
		}
	}
	return dst, nil
}

// Get 根据名称，查找算法
func (inst *ManagerImpl) Get(algorithm string, selector func(reg *keys.Registration) bool) (keys.Algorithm, error) {
	if selector == nil {
		selector = inst.myDefaultSelector
	}
	algorithm = inst.normalizeName(algorithm)
	all := inst.getAll()
	for _, item1 := range all {
		if item1.Name != algorithm {
			continue
		}
		item2 := &keys.Registration{}
		*item2 = *item1
		if selector(item2) {
			return item1.Algorithm, nil
		}
	}
	return nil, fmt.Errorf("no algorithm with name:" + algorithm)
}

func (inst *ManagerImpl) myDefaultSelector(reg *keys.Registration) bool {
	return true
}

func (inst *ManagerImpl) sort(src []*keys.Registration) []*keys.Registration {
	s := &ManagerImpl{items: src}
	sort.Sort(s)
	return s.items
}

func (inst *ManagerImpl) Len() int {
	return len(inst.items)
}

func (inst *ManagerImpl) Less(a, b int) bool {
	p1 := inst.items[a].Priority
	p2 := inst.items[b].Priority
	return p1 > p2
}

func (inst *ManagerImpl) Swap(a, b int) {
	list := inst.items
	list[a], list[b] = list[b], list[a]
}
