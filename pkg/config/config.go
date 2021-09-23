package config

import "github.com/spf13/viper"

type ConfigValues string

const (
	DefaultMiddleware ConfigValues = "default.middleware"
)

func init() {
	setDefaults()
	viper.SetEnvPrefix("easytunnel")

}

func setDefaults() {
	viper.SetDefault("default.middleware", "ssh_docker")
}

func GetStringValue(key ConfigValues) string {
	return viper.GetString(string(key))
}
