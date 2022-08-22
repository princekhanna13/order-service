//Package config provides the available configurations of the project, and access to them.
package config

import "github.com/vrischmann/envconfig"

var appConfig Config

type Config struct {
	Rest struct {
		Port       string `envconfig:"default=8080,optional"`
		APIVersion string `envconfig:"default=v1,optional"`
		PathPrefix string `envconfig:"default=/gd,optional"`
	}
	LogLevel string `envconfig:"default=INFO,optional"`
}

func InitConfig() error {
	appConfig = Config{}
	err := envconfig.Init(&appConfig)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig() Config {
	return appConfig
}
