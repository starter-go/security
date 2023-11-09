package internal

import "github.com/starter-go/base/lang"

type maxAgeComputer struct {

	// input

	ageMax     lang.Seconds
	ageMin     lang.Seconds
	ageDefault lang.Seconds
}

// compute
// 输入：(createdAt:创建时间戳, maxAge:保质期, expiredAt:过期时间戳)
// 输出：(create2:创建时间戳, maxAge2:保质期, exp2:过期时间戳)
func (inst *maxAgeComputer) compute(createdAt lang.Time, maxAge lang.Seconds, expiredAt lang.Time) (lang.Time, lang.Seconds, lang.Time) {

	// init
	now := lang.Now()
	if createdAt == 0 {
		createdAt = now
	}
	if maxAge == 0 {
		if expiredAt == 0 {
			maxAge = inst.ageDefault
		} else {
			du := expiredAt.Sub(createdAt)
			maxAge = lang.NewSeconds(du)
		}
	}

	// check range
	if maxAge < inst.ageMin {
		maxAge = inst.ageMin
	}
	if maxAge > inst.ageMax {
		maxAge = inst.ageMax
	}

	// remake
	expiredAt = createdAt.Add(maxAge.Duration())

	return createdAt, maxAge, expiredAt
}
