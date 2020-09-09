package entity

import (
	"github.com/jinzhu/gorm"
)

type Weather struct {
	gorm.Model

	WeatherId   int
	Main        string
	Description string
	HourlyWeatherInfoId int
}

func NewWeather(id int, main string, description string,hid int) *Weather {
	return &Weather{
		WeatherId:   id,
		Main:        main,
		Description: description,
		HourlyWeatherInfoId: hid,
	}
}

func (w *Weather) Create(db *gorm.DB) error {
	return db.Create(w).Error
}
