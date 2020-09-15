package api

import (
	"bytes"
	"encoding/csv"
	"github.com/labstack/echo"
	"github.com/nachiguro1003/otenki/src/servicer/app/entity"
	"github.com/nachiguro1003/otenki/src/servicer/frame"
	"net/http"
	"strconv"
	"time"
)

type HourlyWeatherInfo struct {
	entity.HourlyWeatherInfo
	Weathers []entity.Weather `gorm:"ForeignKey:HourlyWeatherInfoId;AssociationForeignKey:ID"`
}

func getWeatherInfo(ot *frame.OtenkiFrame,from,to int32) ([]HourlyWeatherInfo,error) {
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

func GetWeatherInfoHandler(c echo.Context,ot *frame.OtenkiFrame) error {
	from,_ := strconv.Atoi(c.QueryParam("from"))
	to,_ := strconv.Atoi(c.QueryParam("to"))

	hw,err := getWeatherInfo(ot,int32(from),int32(to))
	if err != nil {
		return err
	}
	res := bytes.Buffer{}
	w := csv.NewWriter(&res)
	w.Write([]string{"Date", "Temperature", "WeatherId", "Weather", "Description"})

	for _,v := range hw {
		for _,l := range v.Weathers {
			date:= time.Unix(int64(v.Date),0)
			temp:= strconv.Itoa(int(v.Temperature))
			wid:= strconv.Itoa(l.WeatherId)
			w.Write([]string{date.Format(time.RFC3339),temp,wid,l.Main,l.Description})
		}
	}
	w.Flush()

	return c.Blob(http.StatusOK,"text/csv",res.Bytes())
}
