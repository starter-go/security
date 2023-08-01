package internal

import (
	"crypto/rand"
	"io"
	"sync"

	"github.com/starter-go/security/random"
)

// DefaultRandomSource ... 默认随机源
type DefaultRandomSource struct {

	//starter:component
	_as func(random.ProviderRegistry) //starter:as(".")

	mutex sync.Mutex
	r     io.Reader
}

func (inst *DefaultRandomSource) _impl() (random.Source, random.Provider, random.ProviderRegistry, io.Reader) {
	return inst, inst, inst, inst
}

func (inst *DefaultRandomSource) Reader() io.Reader {
	return inst
}

func (inst *DefaultRandomSource) Source() random.Source {
	return inst
}

func (inst *DefaultRandomSource) Registration() *random.ProviderRegistration {
	return &random.ProviderRegistration{
		Name:     "default",
		Enabled:  true,
		Provider: inst,
	}
}

func (inst *DefaultRandomSource) Read(p []byte) (n int, err error) {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	r := inst.r
	if r == nil {
		r = rand.Reader
		inst.r = r
	}

	return r.Read(p)
}
