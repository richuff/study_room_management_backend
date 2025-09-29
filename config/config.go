package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App struct {
		Mode string `yaml:"mode"`
		Port int    `yaml:"port"`
	} `yaml:"app"`
	Storage struct {
		Local localConf `yaml:"local"`
	} `yaml:"storage"`
}

type localConf struct {
	Path string `yaml:"path"`
}

var C = read()

func read() *Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
