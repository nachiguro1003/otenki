package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

func LoadConfigForYaml() (*Config, error) {
	buf, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		log.Print(err)
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	return &cfg, err
}
