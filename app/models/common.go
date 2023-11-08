package models

import (
	"gorm.io/gorm"
	"time"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primarykey;AUTO_INCREMENT"`
}

// 创建、更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
}

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
