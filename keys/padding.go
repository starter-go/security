package keys

import (
	"bytes"
	"fmt"
)

// Padding 是提供填充方法的接口
type Padding interface {

	// 取模式名称
	Mode() PaddingMode

	// 填充
	Pad(src []byte, blockSize int) ([]byte, error)

	// 去除填充
	Unpad(src []byte, blockSize int) ([]byte, error)
}

////////////////////////////////////////////////////////////////////////////////

// PaddingCache 是一个简单的填充模式缓存
type PaddingCache struct {
	p Padding
}

// Get 获取指定的填充模式
func (inst *PaddingCache) Get(want PaddingMode) (Padding, error) {
	p := inst.p
	if p != nil {
		if p.Mode() == want {
			return p, nil
		}
	}
	p, err := GetPadding(want)
	if err != nil {
		return nil, err
	}
	inst.p = p
	return p, nil
}

////////////////////////////////////////////////////////////////////////////////

// GetPadding 获取指定的填充模式
func GetPadding(mode PaddingMode) (Padding, error) {
	switch mode {
	case PKCS5Padding:
		return createPkcs5padding(mode)
	case PKCS7Padding:
		return createPkcs7padding(mode)
	}
	return nil, fmt.Errorf("unsupported padding mode:%s", mode)
}

////////////////////////////////////////////////////////////////////////////////

type pkcs7padding struct {
	mode PaddingMode
}

func (inst *pkcs7padding) Mode() PaddingMode {
	return inst.mode
}

func (inst *pkcs7padding) Pad(src []byte, blockSize int) ([]byte, error) {
	err := inst.checkBlockSize(blockSize)
	if err != nil {
		return nil, err
	}
	olderSize := len(src)
	remainder := (olderSize % blockSize)
	padSize := blockSize - remainder
	padData := make([]byte, padSize)
	for i := 0; i < padSize; i++ {
		padData[i] = byte(padSize)
	}
	buffer := bytes.NewBuffer(src)
	buffer.Write(padData)
	return buffer.Bytes(), nil
}

func (inst *pkcs7padding) Unpad(src []byte, blockSize int) ([]byte, error) {
	err := inst.checkBlockSize(blockSize)
	if err != nil {
		return nil, err
	}
	srcSize := len(src)
	if srcSize < 1 {
		return nil, fmt.Errorf("bad raw data size:%d", srcSize)
	}
	padSize := int(src[srcSize-1])
	if srcSize < padSize || blockSize < padSize {
		return nil, fmt.Errorf("bad padding size:%d", padSize)
	}
	return src[0 : srcSize-padSize], nil
}

func (inst *pkcs7padding) isPKCS5() bool {
	return inst.mode == PKCS5Padding
}

func (inst *pkcs7padding) checkBlockSize(blockSize int) error {
	if inst.isPKCS5() {
		if blockSize == 8 {
			return nil
		}
	} else {
		if 1 <= blockSize && blockSize <= 255 {
			return nil
		}
	}
	return fmt.Errorf("bad block size:%d as mode:%s", blockSize, inst.mode)
}

func createPkcs7padding(mode PaddingMode) (Padding, error) {
	mode = PKCS7Padding
	p := &pkcs7padding{mode}
	return p, nil
}

func createPkcs5padding(mode PaddingMode) (Padding, error) {
	mode = PKCS5Padding
	p := &pkcs7padding{mode}
	return p, nil
}

////////////////////////////////////////////////////////////////////////////////
