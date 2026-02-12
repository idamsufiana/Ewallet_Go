package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")        // root
	viper.AddConfigPath("./config") // folder config

	viper.AutomaticEnv() // allow ENV override

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}
