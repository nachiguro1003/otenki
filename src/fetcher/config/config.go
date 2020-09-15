package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

func LoadConfigForYaml() (*Config, error) {
	filename := os.Args[1]
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	return &cfg, err
}

