package main

import (
	"github.com/drivernation/kaiju"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	kaiju.Config `yaml:",inline"`
	Saying       string `yaml:"saying" json:"saying"`
}

// Load yaml configuration from file. Returns a populated Config object and and optional error if something went wrong.
func LoadConfigYaml(configFile string) (Config, error) {
	c := Config{
		Config: kaiju.Config{
			BindHost: "localhost",
			Port:     8080,
		},
	}
	b, err := ioutil.ReadFile(configFile)

	if err == nil {
		err = yaml.Unmarshal(b, &c)
	}

	return c, err
}
