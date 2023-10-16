package internal

import (
	"testing"

	"github.com/starter-go/security/random"
)

func TestUUIDService(t *testing.T) {

	var ser random.UUIDService
	ser = &UUIDServiceImpl{}

	u1 := ser.Build().Generate()
	u2 := ser.Build().Generate()
	u3 := ser.Build().Class("foo").ID("bar").Generate()

	t.Log("uuid_1:" + u1)
	t.Log("uuid_2:" + u2)
	t.Log("uuid_3:" + u3)
}
