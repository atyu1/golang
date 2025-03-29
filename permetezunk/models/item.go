package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	Title       string `gorm:"type:text"`
	Description string `gorm:"type:text"`
}
