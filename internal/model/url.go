package model

import (
	"time"
	"gorm.io/gorm"
)

type URL struct{
	ID uint `gorm:"primaryKey"`
	ShortCode string `gorm:"uniqueIndex;not null"`
	OriginalURL string `gorm:"not null"`
	CreatedAt time.Time
	Clicks int `gorm:"default:0"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}