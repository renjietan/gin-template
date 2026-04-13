package entity

import (
	"time"
)

type BaseEntity struct {
	// 主键，类型为 bigint unsigned，自增
	ID uint `gorm:"primarykey;autoIncrement" json:"id"`
	// 时间字段，使用 datetime(3) 以支持毫秒精度
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`
	// 软删除字段，使用指针，允许为 NULL
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime(3);index" json:"deleted_at,omitempty"`
}
