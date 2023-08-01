package random

import "io"

// Source 随机数来源
type Source interface {
	Reader() io.Reader
}

// Service 随机数服务
type Service interface {
	Source() Source
	Reader() io.Reader

	NextString(length int) string
	NextBytes(length int) []byte
	NextInt() int
	NextInt64() int64
}
