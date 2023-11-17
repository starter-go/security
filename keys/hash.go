package keys

import (
	"crypto"
	"fmt"
	"hash"
)

// GetHash 获取指定名称的 hash 算法
func GetHash(algorithm crypto.Hash) (hash.Hash, error) {
	if algorithm.Available() {
		return algorithm.New(), nil
	}
	name := algorithm.String()
	return nil, fmt.Errorf("unsupported hash algorithm: %s", name)
}
