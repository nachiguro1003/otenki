package config

type Config struct {
	OpenWeatherApi openWeatherApiConfig `yaml:"open_weather_api"`
	Postgres       postgres             `yaml:"postgres"`
}

type openWeatherApiConfig struct {
	Lang      string `yaml:"lang"`
	ApiKey    string `yaml:"api_key"`
	Endpoint  string `yaml:"endpoint"`
	Latitude  string `yaml:"latitude"`
	Longitude string `yaml:"longitude"`
}

type postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"db_name"`
	Password string `yaml:"password"`
}
