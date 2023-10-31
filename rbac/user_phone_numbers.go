package rbac

import (
	"context"
	"fmt"
	"strings"
)

// PhoneNumberID ...
type PhoneNumberID int64

// PhoneNumber 表示电话号码
type PhoneNumber string

// SimplePhoneNumber 表示简短的电话号码, 标准化的取值为纯数字形式，不含任何其它字符，like:"12345678901"
type SimplePhoneNumber PhoneNumber

// FullPhoneNumber 表示完整的电话号码, like: "+86-123-4567-8901"
type FullPhoneNumber PhoneNumber

// PurePhoneNumber 表示完整(且纯粹)的电话号码, like: "8612345678901"
type PurePhoneNumber PhoneNumber

// PhoneNumberDTO ...
type PhoneNumberDTO struct {
	ID PhoneNumberID `json:"id"`

	BaseDTO

	RegionCode2  RegionPhoneCode   `json:"region"`
	SimpleNumber SimplePhoneNumber `json:"simple_number"`
	FullNumber   FullPhoneNumber   `json:"full_number"`
}

// PhoneNumberQuery 查询参数
type PhoneNumberQuery struct {
	Conditions Conditions
	Pagination Pagination
	All        bool // 查询全部条目
}

// PhoneNumberService ...
type PhoneNumberService interface {
	Insert(c context.Context, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Update(c context.Context, id PhoneNumberID, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Delete(c context.Context, id PhoneNumberID) error

	Find(c context.Context, id PhoneNumberID) (*PhoneNumberDTO, error)
	List(c context.Context, q *PhoneNumberQuery) ([]*PhoneNumberDTO, error)
}

////////////////////////////////////////////////////////////////////////////////

func (num PhoneNumber) String() string {
	return string(num)
}

// // Normalize 标准化
// func (num PhoneNumber) Normalize() PhoneNumber {
// 	return string(num)
// }

////////////////////////////////////////////////////////////////////////////////

func (num SimplePhoneNumber) String() string {
	return string(num)
}

// Normalize 标准化
func (num SimplePhoneNumber) Normalize() SimplePhoneNumber {
	str := num.String()
	str = normalizePureNumber(str)
	return SimplePhoneNumber(str)
}

////////////////////////////////////////////////////////////////////////////////

// ParseFullPhoneNumber 把输入的字符串解析为完整的电话号码
func ParseFullPhoneNumber(str string) (FullPhoneNumber, error) {
	region, phone, err := parseFullPhoneNumberParts(str)
	if err != nil {
		return "", err
	}
	s2 := "+" + region.String() + "-" + phone.String()
	return FullPhoneNumber(s2), nil
}

func (num FullPhoneNumber) String() string {
	return string(num)
}

// Pure 转换为纯号码形式
func (num FullPhoneNumber) Pure() PurePhoneNumber {
	p := PurePhoneNumber(num)
	return p.Normalize()
}

// Parts 拆解为：区号 + SimplePhoneNumber
func (num FullPhoneNumber) Parts() (RegionPhoneCode, SimplePhoneNumber, error) {
	str := num.String()
	return parseFullPhoneNumberParts(str)
}

// Normalize 标准化
func (num FullPhoneNumber) Normalize() FullPhoneNumber {
	str := num.String()
	num2, err := ParseFullPhoneNumber(str)
	if err == nil {
		return num2
	}
	return num
}

////////////////////////////////////////////////////////////////////////////////

func (num PurePhoneNumber) String() string {
	return string(num)
}

// Normalize 标准化
func (num PurePhoneNumber) Normalize() PurePhoneNumber {
	str := num.String()
	str = normalizePureNumber(str)
	return PurePhoneNumber(str)
}

////////////////////////////////////////////////////////////////////////////////
// internal functions

func normalizePureNumber(src string) string {
	chs := []rune(src)
	b := strings.Builder{}
	for _, ch := range chs {
		if '0' <= ch && ch <= '9' {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func parseFullPhoneNumberParts(str string) (RegionPhoneCode, SimplePhoneNumber, error) {
	b := strings.Builder{}
	src := []rune(str)
	countPlus := 0
	countMin := 0
	part1 := ""
	part2 := ""
	for _, ch := range src {
		if '0' <= ch && ch <= '9' {
			b.WriteRune(ch)
		} else if ch == '+' {
			countPlus++
			if countPlus == 1 && countMin == 0 {
				b.Reset()
			} else {
				return "", "", fmt.Errorf("bad full-phone-number: " + str)
			}
		} else if ch == '-' {
			countMin++
			if countPlus == 1 && countMin == 1 {
				part1 = b.String()
				b.Reset()
			}
		} else {
			// nop
		}
	}
	part2 = b.String()
	if part1 == "" || part2 == "" {
		return "", "", fmt.Errorf("bad full-phone-number: " + str)
	}
	return RegionPhoneCode(part1), SimplePhoneNumber(part2), nil
}
