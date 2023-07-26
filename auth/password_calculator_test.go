package auth

import (
	"testing"

	"github.com/starter-go/base/lang"
)

func TestPasswordCalc(t *testing.T) {

	pc := &PasswordCalculator{}

	// step-1: empty
	pc.Init(nil, nil)
	err := pc.Verify(nil)
	if err != nil {
		t.Log(err)
	}

	// step-2: normal data
	salt := lang.Hex("abc0123456789def")
	want := lang.Hex("fc20c12412e4780595727a6e934e62b1ccfd37e0609e43b204c04d580d5e63bc")
	plain := "hello"

	pc.Init(want.Bytes(), salt.Bytes())
	err = pc.Verify([]byte(plain))
	if err != nil {
		t.Log(err)
	}

}
