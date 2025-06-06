// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameIdentityCard = "identity_card"

// IdentityCard 身份证
type IdentityCard struct {
	ID        int                   `gorm:"column:id;primaryKey;autoIncrement:true;comment:记录ID" json:"id"`                                                          // 记录ID
	Number    string                `gorm:"column:number;not null;comment:身份证号码" json:"number"`                                                                      // 身份证号码
	UserID    int                   `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                                                     // 用户ID
	CreatedAt time.Time             `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建日期" json:"created_at,omitzero" time_format:"sql_datetime"` // 创建日期
	UpdatedAt time.Time             `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新日期" json:"updated_at,omitzero" time_format:"sql_datetime"` // 更新日期
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;not null;comment:删除时间戳[0:未删除,非0:删除时间戳]" json:"-"`                                                       // 删除时间戳[0:未删除,非0:删除时间戳]
}

// TableName IdentityCard's table name
func (*IdentityCard) TableName() string {
	return TableNameIdentityCard
}
