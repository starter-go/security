package rbac

import "context"

// RegionID ...
type RegionID int

// RegionPhoneCode 是数字形式的国际电话区号，
// 例如：中国(+86)， 法国(+33)， 俄国(+7)， 美国(+1)， 英国(+44)
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

	SimpleName string          // 区域简称，例如：chn(中国), fra(France), usa(United States)
	Code2      RegionCode2     // 二字符区域码
	Code3      RegionCode3     // 三字符区域码
	PhoneCode  RegionPhoneCode // 电话区域码
}

// RegionQuery 查询参数
type RegionQuery struct {
	Pagination Pagination
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

func (code RegionCode2) String() string {
	return string(code)
}

func (code RegionCode3) String() string {
	return string(code)
}
