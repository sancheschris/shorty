package repository

import (
	"github.com/sancheschris/shorty/internal/database"
	"github.com/sancheschris/shorty/internal/model"
	"gorm.io/gorm"
)


func SaveShortURL(url *model.URL) error {
	return database.DB.Create(url).Error
}

func GetURLByShortCode(code string) (*model.URL, error) {
	var url model.URL
	err := database.DB.Where("short_code = ?", code).First(&url).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &url, err
}

func IncrementClicks(code string) error {
    return database.DB.Model(&model.URL{}).
        Where("short_code = ?", code).
        UpdateColumn("clicks", gorm.Expr("clicks + 1")).Error
}