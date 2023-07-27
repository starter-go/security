package rbac

import (
	"github.com/starter-go/base/lang"
	"gorm.io/gorm"
)

// ConvertTimeToGorm ...
func ConvertTimeToGorm(t lang.Time) gorm.DeletedAt {
	dst := gorm.DeletedAt{}
	dst.Valid = (t > 0)
	dst.Time = t.Time()
	return dst
}

// ConvertTimeFromGorm ...
func ConvertTimeFromGorm(t gorm.DeletedAt) lang.Time {
	if t.Valid {
		return lang.NewTime(t.Time)
	}
	return 0
}

// CopyBaseFieldsFromEntityToDTO ...
func CopyBaseFieldsFromEntityToDTO(src *BaseEntity, dst *BaseDTO) {

	dst.UUID = src.UUID

	dst.CreatedAt = lang.NewTime(src.CreatedAt)
	dst.UpdatedAt = lang.NewTime(src.UpdatedAt)
	dst.DeletedAt = ConvertTimeFromGorm(src.DeletedAt)

	dst.Owner = src.Owner
	dst.Updater = src.Updater
	dst.Creator = src.Creator
}

// CopyBaseFieldsFromDtoToEntity ...
func CopyBaseFieldsFromDtoToEntity(src *BaseDTO, dst *BaseEntity) {

	dst.UUID = src.UUID

	// 这3个字段由 gorm 自动设置
	// dst.CreatedAt = src.CreatedAt.Time()
	// dst.UpdatedAt = src.UpdatedAt.Time()
	// dst.DeletedAt = ConvertTimeToGorm(src.DeletedAt)

	dst.Owner = src.Owner
	dst.Updater = src.Updater
	dst.Creator = src.Creator
}
