package frame

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nachiguro1003/otenki/src/fetcher/config"
	"github.com/nachiguro1003/otenki/src/fetcher/db"
)

type OtenkiFrame struct {
	Config *config.Config
	PostgresDB *db.PostgresEntity
	Echo *echo.Echo
}

func InitOtenkiFrame() (*OtenkiFrame, error){
	cfg, err := config.LoadConfigForYaml()
	if err != nil {
		return nil,err
	}
	d := db.InitDB(
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.DBName,
		cfg.Postgres.Password)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &OtenkiFrame{
		Config: cfg,
		PostgresDB: d,
		Echo: e,
	},nil
}
