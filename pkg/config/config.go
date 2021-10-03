package config

import (
	log2 "easytunnel/pkg/log"
	"easytunnel/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type ConfigValues string

const (
	DefaultMiddleware ConfigValues = "default.middleware"
	RemoteHost        ConfigValues = "RemoteHost"
	RemotePort        ConfigValues = "RemotePort"
)

var (
	log      logrus.Logger = *log2.GetLogger()
	ClientID string
)

func init() {
	setDefaults()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType(".yaml")
	viper.SetEnvPrefix("easytunnel")
	getOrSetClientID()

	err := viper.SafeWriteConfig()
	if err != nil {
		panic(err)
	}

}

func getOrSetClientID() {
	clientIDConfig := viper.New()
	clientIDConfig.SetConfigFile("/easytunnel/config/client_id.yaml")
	//clientIDConfig.AddConfigPath("./config/")
	if !utils.FileExits("./config") {
		err := os.Mkdir("./config", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	clientIDConfigReadErr := clientIDConfig.ReadInConfig()
	if clientIDConfigReadErr != nil {
		log.Error(clientIDConfigReadErr)
		ClientID, _ = gonanoid.ID(20)
		clientIDConfig.Set("id", ClientID)
		err := clientIDConfig.WriteConfigAs("./config/client_id.yaml")
		if err != nil {
			log.Error(err)
			panic(err)
		}
	}
	ClientID = clientIDConfig.GetString("id")
}

func setDefaults() {
	viper.SetDefault(string(DefaultMiddleware), "ssh_docker")
	viper.SetDefault(string(RemoteHost), "127.0.0.1")
	viper.SetDefault(string(RemotePort), 80)

}

func GetStringValue(key ConfigValues) string {
	return viper.GetString(string(key))
}
