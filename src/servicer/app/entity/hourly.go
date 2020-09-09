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

