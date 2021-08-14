package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Mode       string
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	DBUrl      string `mapstructure:"DATABASE_URL"`
	//DBHost     string `mapstructure:"DATABASE_HOST"`
	//DBPort     string `mapstructure:"DATABASE_PORT"`
	//DBUser     string `mapstructure:"DATABASE_USER"`
	//DBPass     string `mapstructure:"DATABASE_PASSWORD"`
	//DBName     string `mapstructure:"DATABASE_NAME"`
	//DBSslMode  string `mapstructure:"DATABASE_SSLMODE"`
	//DBTz       string `mapstructure:"DATABASE_TZ"`
}

func (c *Config) checkConfig() error {

	e := reflect.ValueOf(c).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()

		if e.Field(i).Len() == 0 {
			return errors.New(fmt.Sprintf("Config not valid, param empty %v %v\n", varName, varValue))
			//fmt.Printf("Config not valid, param empty %v %v\n", varName, varValue)
		}
	}

	return nil
}

func New() (c *Config, err error) {

	c = &Config{Mode: "prod"}

	for _, arg := range os.Args {
		fmt.Println(arg)
		if strings.HasPrefix(arg, "-test.v=") {
			c.Mode = "test"
		}
		if strings.HasPrefix(arg, "-mode=dev") {
			c.Mode = "dev"
		}
	}

	path := "./"
	if c.Mode == "test" {
		path = "./.."
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("."+c.Mode)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	if err = viper.Unmarshal(&c); err != nil {
		return c, err
	}

	if err = c.checkConfig(); err != nil {
		return c, err
	}

	return c, nil
}
