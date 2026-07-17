package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 模型基类（类似 PHP 的 Model 基类）
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
