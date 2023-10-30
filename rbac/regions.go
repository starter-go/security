package rbac

import (
	"context"
	"strings"
)

// RegionID ...
type RegionID int

// RegionPhoneCode 是数字形式的国际电话区号，
// 例如：中国(86)， 法国(33)， 俄国(7)， 美国(1)， 英国(44)
type RegionPhoneCode string

// RegionCode2 是 ISO 3166-1 标准的二字节地区码
// 例如：中国(CN)， 法国(FR)， 俄国(RU)， 美国(US)， 英国(GB)
type RegionCode2 string

// RegionCode3 是 ISO 3166-1 标准的三字节地区码
// 例如：中国(CHN)， 法国(FRA)， 俄国(RUS)， 美国(USA)， 英国(GBR)
type RegionCode3 string

// RegionDTO 表示国际区号
type RegionDTO struct {
	ID RegionID `json:"id"`

	BaseDTO

	FlagURL     string          `json:"flag_url"`    // 国旗（或区旗）图标的URL
	DisplayName string          `json:"label"`       // 显示名称，通常是本地化的名称
	SimpleName  string          `json:"simple_name"` // 区域简称，例如：chn(中国), fra(France), usa(United States)
	FullName    string          `json:"full_name"`   // 完整的名称，例如：中华人民共和国(PRC)
	Code2       RegionCode2     `json:"code_xx"`     // 二字符区域码
	Code3       RegionCode3     `json:"code_xxx"`    // 三字符区域码
	PhoneCode   RegionPhoneCode `json:"phone_code"`  // 电话区域码
}

// RegionQuery 查询参数
type RegionQuery struct {
	Conditions Conditions
	Pagination Pagination
	All        bool // 查询全部条目
}

// RegionService ...
type RegionService interface {
	Insert(c context.Context, o *RegionDTO) (*RegionDTO, error)
	Update(c context.Context, id RegionID, o *RegionDTO) (*RegionDTO, error)
	Delete(c context.Context, id RegionID) error

	Find(c context.Context, id RegionID) (*RegionDTO, error)
	List(c context.Context, q *RegionQuery) ([]*RegionDTO, error)
}

////////////////////////////////////////////////////////////////////////////////

func (code RegionPhoneCode) String() string {
	return string(code)
}

// Normalize 标准化代码
func (code RegionPhoneCode) Normalize() RegionPhoneCode {
	chs := []rune(code.String())
	b := strings.Builder{}
	for _, ch := range chs {
		if '0' <= ch && ch <= '9' {
			b.WriteRune(ch)
		}
	}
	str := b.String()
	return RegionPhoneCode(str)
}

////////////////////////////////////////////////////////////////////////////////

func (code RegionCode2) String() string {
	return string(code)
}

// Normalize 标准化代码
func (code RegionCode2) Normalize() RegionCode2 {
	str := code.String()
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	return RegionCode2(str)
}

////////////////////////////////////////////////////////////////////////////////

func (code RegionCode3) String() string {
	return string(code)
}

// Normalize 标准化代码
func (code RegionCode3) Normalize() RegionCode3 {
	str := code.String()
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	return RegionCode3(str)
}

////////////////////////////////////////////////////////////////////////////////
