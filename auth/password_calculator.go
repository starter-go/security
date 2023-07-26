package auth

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/vlog"
)

// PasswordCalculator 是默认的密码计算器
type PasswordCalculator struct {
	salt   []byte
	target []byte
}

// Init ...
func (inst *PasswordCalculator) Init(target, salt []byte) {
	inst.salt = salt
	inst.target = target
}

// Compute 计算密码的哈希值
func (inst *PasswordCalculator) Compute(plain []byte) []byte {
	const total = 5
	current := plain
	for step := 0; step < total; step++ {
		if step == 1 {
			current = inst.hashOnce(current, inst.salt)
		} else {
			current = inst.hashOnce(current, nil)
		}
	}
	return current
}

// Verify 验证密码是否正确
func (inst *PasswordCalculator) Verify(plain []byte) error {

	want := inst.target
	have := inst.Compute(plain)
	if bytes.Equal(want, have) {
		return nil
	}

	logger := vlog.GetLogger()
	if logger.IsDebugEnabled() {
		w := lang.HexFromBytes(want)
		h := lang.HexFromBytes(have)
		logger.Debug("user.password.want: %s", w.String())
		logger.Debug("user.password.have: %s", h.String())
	}

	return fmt.Errorf("bad auth")
}

func (inst *PasswordCalculator) hashOnce(b1, b2 []byte) []byte {
	buffer := bytes.NewBuffer(b1)
	buffer.Write(b2)
	data := buffer.Bytes()
	sum := sha256.Sum256(data)
	return sum[:]
}
