package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

func LoadConfigForYaml() (*Config, error) {
	buf, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	return &cfg, err
}
