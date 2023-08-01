package code

import (
	"testing"

	"github.com/starter-go/security/internal"
	"github.com/starter-go/security/random"
	"github.com/starter-go/vlog"
)

func TestRandomCom(t *testing.T) {

	provider := &internal.DefaultRandomSource{}
	ser := &internal.DefaultRandomService{
		Providers: []random.ProviderRegistry{provider},
	}

	tr := &TestRandom{}
	tr.Logger = vlog.GetLogger()
	tr.Rand = ser

	tr.run()
}
