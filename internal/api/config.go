package api

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Name string `yaml:"name"`
	Environment string
	Api struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"api"`
}

// Init is using to initialize the configs
func loadConfigFile() (*Config, error) {

	mode := flag.String("mode", "prod", "mode conf default prod")
	flag.Parse()

	fmt.Println(*mode)


	file := "./config.yaml"
	var apiConfig *Config

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var options map[string]Config
	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return nil, err
	}
	opt := options[*mode]
	opt.Environment = *mode
	apiConfig = &opt

	return apiConfig, nil
}

func NewConfig() (*Config, error)  {
	return loadConfigFile()
}
