package entity

import (
	"github.com/jinzhu/gorm"
)

type HourlyWeatherInfo struct {
	gorm.Model

	Date    int64
	Temperature float32
}

func NewHourly(unixDate int64,temp float32) *HourlyWeatherInfo {
	return &HourlyWeatherInfo{
		Date: unixDate,
		Temperature: temp,
	}
}

func (h *HourlyWeatherInfo) Create(db *gorm.DB) error {
	return db.Create(h).Error
}

func IsExistByDate(db *gorm.DB,unixDate int64) bool {
	return db.
		Model(&HourlyWeatherInfo{}).
		Where("date = ?",unixDate).
		RowsAffected > 0
}
