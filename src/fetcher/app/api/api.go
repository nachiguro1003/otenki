package api

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/otenki/src/fetcher/app/entity"
	"github.com/nachiguro1003/otenki/src/fetcher/frame"
	"io/ioutil"
	"net/http"
)

type HourlyWeatherInfoBinder struct {
	Hourly []HourlyWeatherInfo `json:"hourly"`
}

type HourlyWeatherInfo struct {
	UnixTime    int64     `json:"dt"`
	Temperature float32   `json:"temp"`
	Weather     []weather `json:"weather"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

var (
	client = new(http.Client)
)

func FetchWeatherInfo(ot *frame.OtenkiFrame) error {
	var hw HourlyWeatherInfoBinder

	if err := hw.accessOpenWeatherAPI(ot); err != nil {
		return err
	}

	d, err := ot.PostgresDB.Connect()
	if err != nil {
		return err
	}

	if err = d.Transaction(func(tx *gorm.DB) error {
		v := hw.Hourly[len(hw.Hourly)-1]
		if !entity.IsExistByDate(d, v.UnixTime) {
			he := entity.NewHourly(v.UnixTime, v.Temperature)
			if err := he.Create(d); err != nil {
				return err
			}

			for _, t := range v.Weather {
				we := entity.NewWeather(t.Id, t.Main, t.Description, int(he.ID))
				if err := we.Create(d); err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (hw *HourlyWeatherInfoBinder) accessOpenWeatherAPI(ot *frame.OtenkiFrame) error {
	ow := ot.Config.OpenWeatherApi
	param := fmt.Sprintf("?lat=%s&lon=%s&appid=%s", ow.Latitude, ow.Longitude, ow.ApiKey)
	resp, err := client.Get(ow.Endpoint + param)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &hw); err != nil {
		return err
	}

	return nil
}
