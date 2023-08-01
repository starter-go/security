package internal

import (
	"encoding/base64"
	"fmt"
	"io"

	"github.com/starter-go/security/random"
)

// DefaultRandomService ... 随机数服务
type DefaultRandomService struct {

	//starter:component
	_as func(random.Service) //starter:as("#")

	Providers []random.ProviderRegistry //starter:inject(".")

	cachedSource random.Source
}

func (inst *DefaultRandomService) _impl() random.Service {
	return inst
}

func (inst *DefaultRandomService) loadSource() (random.Source, error) {

	src := inst.Providers
	dst := inst.cachedSource
	for _, r1 := range src {
		if r1 == nil {
			continue
		}
		r2 := r1.Registration()
		if r2 == nil {
			continue
		}
		p := r2.Provider
		if r2.Enabled && p != nil {
			dst = p.Source()
		}
	}

	if dst == nil {
		return nil, fmt.Errorf("no available random.Provider")
	}
	return dst, nil
}

// Source ...
func (inst *DefaultRandomService) Source() random.Source {

	s := inst.cachedSource
	if s != nil {
		return s
	}

	s2, err := inst.loadSource()
	if err != nil {
		panic(err)
	}

	s = s2
	inst.cachedSource = s2
	return s
}

// Reader ...
func (inst *DefaultRandomService) Reader() io.Reader {
	return inst.Source().Reader()
}

// NextString ...
func (inst *DefaultRandomService) NextString(length int) string {
	data := inst.NextBytes(length)
	str := base64.StdEncoding.EncodeToString(data)
	if len(str) > length {
		return str[0:length]
	}
	return str
}

// NextBytes ...
func (inst *DefaultRandomService) NextBytes(length int) []byte {
	buffer := make([]byte, length)
	count, err := inst.Reader().Read(buffer)
	if err != nil {
		panic(err)
	}
	if count != length {
		panic("bad random data size")
	}
	return buffer
}

// NextInt ...
func (inst *DefaultRandomService) NextInt() int {
	n := inst.NextInt64()
	return int(n)
}

// NextInt64 ...
func (inst *DefaultRandomService) NextInt64() int64 {
	n := int64(0)
	buffer := inst.NextBytes(8)
	for _, b := range buffer {
		n = (n << 8) | int64(b&0xff)
	}
	return n
}
