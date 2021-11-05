package config

import (
	"github.com/foxfurry/go_delivery/internal/infrastracture/logger"
	"github.com/spf13/viper"
)

func LoadConfig(){
	viper.AddConfigPath("./config")

	viper.SetConfigName("cfg")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		logger.LogPanicF("Could not read config file: %v", err)
	}
}
