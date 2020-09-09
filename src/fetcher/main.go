package main

import (
	"github.com/labstack/echo"
	"github.com/nachiguro1003/otenki/src/fetcher/app/api"
	"github.com/nachiguro1003/otenki/src/fetcher/frame"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
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
	g := ot.Echo.Group("/batch")

	// routingは少ないためserver開始に含めている。
	// 拡張するのであれば切り出す
	g.GET("/health_check", func(c echo.Context) error {
		return c.JSON(http.StatusOK,"OK")
	})
	g.GET("/hourly_weather", func(c echo.Context) error {
		err := api.FetchWeatherInfo(ot)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK,"Fetching Succeed")
	})

	err := ot.Echo.Start(":8000")
	if err != nil {
		return err
	}

	return nil
}
