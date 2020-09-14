package main

import (
	"bytes"
	"encoding/csv"
	"github.com/labstack/echo"
	"github.com/nachiguro1003/otenki/src/servicer/app/api"
	"log"
	"strconv"
	"time"

	//"github.com/nachiguro1003/otenki/src/servicer/app/api"
	"github.com/nachiguro1003/otenki/src/servicer/frame"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Printf("err = %s",err)
	}
}

func run() error {

	ot,err := frame.InitOtenkiFrame()
	if err != nil {
		return err
	}

	err = serve(ot)
	if err != nil {
		return err
	}

	return nil
}

func serve(ot *frame.OtenkiFrame) error {
	ot.Echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK,"Hello world")
	})

	g := ot.Echo.Group("/batch")

	// routingは少ないためserver開始に含めている。
	// 拡張するのであれば切り出す
	g.GET("/health_check", func(c echo.Context) error {
		return c.JSON(http.StatusOK,"OK")
	})
	g.GET("/hourly_weather", func(c echo.Context) error {
		from,_ := strconv.Atoi(c.QueryParam("from"))
		to,_ := strconv.Atoi(c.QueryParam("to"))

		hw,err := api.GetWeatherInfo(ot,int32(from),int32(to))
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
	})

	err := ot.Echo.Start(":8001")
	if err != nil {
		return err
	}

	return nil
}
