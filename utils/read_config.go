package utils

import (
	"api/config"
	"log"

	"github.com/spf13/viper"
)

func ReadEnv() (cfg config.Config) {
	viper.SetConfigFile("config.dev.yml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading env: %s\n", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error reading env: %s\n", err)
	}
	return
}
