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
func (inst *maxAgeComputer) compute(createdAt lang.Time, maxAge lang.Milliseconds, expiredAt lang.Time) (lang.Time, lang.Milliseconds, lang.Time) {

	max := inst.sec2ms(inst.ageMax)
	min := inst.sec2ms(inst.ageMin)
	def := inst.sec2ms(inst.ageDefault)

	// init
	now := lang.Now()
	if createdAt == 0 {
		createdAt = now
	}
	if maxAge == 0 {
		if expiredAt == 0 {
			maxAge = def
		} else {
			du := expiredAt.Sub(createdAt)
			maxAge = lang.NewMilliseconds(du)
		}
	}

	// check range
	if maxAge < min {
		maxAge = min
	}
	if maxAge > max {
		maxAge = max
	}

	// remake
	expiredAt = createdAt.Add(maxAge.Duration())

	return createdAt, maxAge, expiredAt
}

func (inst *maxAgeComputer) sec2ms(sec lang.Seconds) lang.Milliseconds {
	du := sec.Duration()
	return lang.NewMilliseconds(du)
}
