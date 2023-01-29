package utils

import (
	"api/config"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReadEnv() (cfg config.Config) {
	viper.SetConfigFile("config.dev.yml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Errorf("Error reading env: %s\n", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Errorf("Error reading env: %s\n", err)
		os.Exit(1)
	}
	return
}
