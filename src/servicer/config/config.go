package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfigForYaml() (*Config, error) {
	filename := os.Args[1]
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	return &cfg, err
}
