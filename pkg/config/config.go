package config

import "github.com/spf13/viper"

type Config struct {
	DefaultMiddleware string
}

var config Config

func Init() {
	viper.SetEnvPrefix("easytunnel")

}

func GetConfig() Config {
	return config
}
