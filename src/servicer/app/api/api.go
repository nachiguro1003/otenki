package api

import (
	"github.com/nachiguro1003/otenki/src/servicer/app/entity"
	"github.com/nachiguro1003/otenki/src/servicer/frame"
)

type HourlyWeatherInfo struct {
	entity.HourlyWeatherInfo
	Weathers []entity.Weather `gorm:"ForeignKey:HourlyWeatherInfoId;AssociationForeignKey:ID"`
}

func GetWeatherInfo(ot *frame.OtenkiFrame,from,to int32) ([]HourlyWeatherInfo,error) {
	var hw []HourlyWeatherInfo

	d, err := ot.PostgresDB.Connect()
	if err != nil {
		return nil,err
	}

	err = d.Where("date between ? and ?",from,to).
		Preload("Weathers").
		Find(&hw).
		Error
	if err != nil {
		return nil,err
	}

	return hw,nil
}
