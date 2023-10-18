package internal

import (
	"testing"

	"github.com/starter-go/security/auth"
)

func TestAuthRegSorter(t *testing.T) {

	list := make([]*auth.Registration, 0)

	list = append(list, &auth.Registration{Priority: 2})
	list = append(list, &auth.Registration{Priority: 0})
	list = append(list, &auth.Registration{Priority: 10})
	list = append(list, &auth.Registration{Priority: 18})
	list = append(list, &auth.Registration{Priority: 2})
	list = append(list, &auth.Registration{Priority: 3})

	s := authRegistrationSorter{}
	s.sort(list)
	for _, r := range list {
		t.Logf("auth.Registration with Priority: %d", r.Priority)
	}
	t.Log("done")
}
