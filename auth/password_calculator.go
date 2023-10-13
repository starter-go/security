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

// Reset ...
func (inst *PasswordCalculator) Reset() *PasswordCalculator {
	inst.salt = nil
	inst.target = nil
	return inst
}

// WithSalt ...
func (inst *PasswordCalculator) WithSalt(salt []byte) *PasswordCalculator {
	inst.salt = salt
	return inst
}

// WithTarget ...
func (inst *PasswordCalculator) WithTarget(target []byte) *PasswordCalculator {
	inst.target = target
	return inst
}

// Compute 计算密码的哈希值
func (inst *PasswordCalculator) Compute(plain []byte) []byte {
	const total = 5
	current := plain
	for i := 0; i < total; i++ {
		if i == 1 {
			current = inst.step(current, inst.salt)
		} else {
			current = inst.step(current, nil)
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

func (inst *PasswordCalculator) step(b1, b2 []byte) []byte {
	buffer := &bytes.Buffer{}
	buffer.Write(b1)
	buffer.Write(b2)
	data := buffer.Bytes()
	sum := sha256.Sum256(data)
	return sum[:]
}
