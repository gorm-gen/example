// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameLanguage = "language"

// Language 语言表
type Language struct {
	ID        int                   `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户ID" json:"id"`                                                          // 用户ID
	Name      string                `gorm:"column:name;not null;comment:语言名称" json:"name"`                                                                           // 语言名称
	CreatedAt time.Time             `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建日期" json:"created_at,omitzero" time_format:"sql_datetime"` // 创建日期
	UpdatedAt time.Time             `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新日期" json:"updated_at,omitzero" time_format:"sql_datetime"` // 更新日期
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;not null;comment:删除时间戳[0:未删除,非0:删除时间戳]" json:"-"`                                                       // 删除时间戳[0:未删除,非0:删除时间戳]
	Users     []User                `gorm:"foreignKey:ID;joinForeignKey:LanguageID;joinReferences:UserID;many2many:user_language;references:ID" json:"users,omitempty"`
}

// TableName Language's table name
func (*Language) TableName() string {
	return TableNameLanguage
}
