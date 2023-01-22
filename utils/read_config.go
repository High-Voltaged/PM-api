package utils

import (
	"api/config"
	"log"
	"os"

	"github.com/spf13/viper"
)

func ReadEnv() (cfg config.Config) {
	viper.SetConfigFile("config.dev.yml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading env: %s\n", err)
		os.Exit(2)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error reading env: %s\n", err)
		os.Exit(2)
	}
	return
}
