package internal

import (
	"sort"

	"github.com/starter-go/security/auth"
)

type authRegistrationSorter struct {
	list []*auth.Registration
}

func (inst *authRegistrationSorter) sort(list []*auth.Registration) {
	inst.list = list
	sort.Sort(inst)
}

func (inst *authRegistrationSorter) Len() int {
	return len(inst.list)
}

func (inst *authRegistrationSorter) Less(a, b int) bool {
	ia := inst.list[a]
	ib := inst.list[b]
	return ia.Priority > ib.Priority // 优先级高的，排前面
}

func (inst *authRegistrationSorter) Swap(a, b int) {
	l := inst.list
	l[a], l[b] = l[b], l[a]
}
