package service

import (
	"github.com/sancheschris/shorty/internal/model"
	"github.com/sancheschris/shorty/internal/repository"
	"github.com/sancheschris/shorty/internal/util"
)


func CreateShortURL(originalURL string) (*model.URL, error) {
	shortCode := util.GenerateShortCode()

	url := &model.URL{
		ShortCode: shortCode,
		OriginalURL: originalURL,
	}

	err := repository.SaveShortURL(url)
	if err != nil {
		return nil, err
	}

	return url, nil
}