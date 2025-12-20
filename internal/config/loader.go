package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var AppCfgs *AppConfigs
func Load() {
	configPath := "./internal/config/config.yml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}

	viper.SetConfigFile(configPath)

	if err:=viper.ReadInConfig();err!=nil{
		log.Println("Viper config load error: ",err.Error())
		return
	}

	AppCfgs=&AppConfigs{}
	if err:=viper.Unmarshal(AppCfgs);err!=nil{
		log.Println("Config unmarshal error: ",err.Error())
		return
	}
	log.Println(AppCfgs)
}