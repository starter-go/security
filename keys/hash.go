package keys

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"strings"
)

// GetHash 获取指定名称的 hash 算法
func GetHash(algorithm string) (hash.Hash, error) {

	name := algorithm
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)

	switch name {
	case "sha1":
		return sha1.New(), nil
	case "sha256":
		return sha256.New(), nil
	case "sha512":
		return sha512.New(), nil
	case "md5":
		return md5.New(), nil
	}
	return nil, fmt.Errorf("unsupported hash algorithm: %s", name)
}
