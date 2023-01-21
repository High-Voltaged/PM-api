package main

import (
	"api/config"
	db "api/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func initRouter(cfg *config.Config) {
	r := mux.NewRouter()

	srv := cfg.Server
	srvConfig := fmt.Sprintf("%s:%s", srv.Host, srv.Port)

	err := http.ListenAndServe(srvConfig, r)
	if err != nil {
		log.Fatalf("Error starting the server: %s\n", err)
		os.Exit(1)
	} else {
		log.Printf("Listening on port %s", srv.Port)
	}
}

func readEnv() (cfg config.Config) {
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

func main() {
	var cfg = readEnv()
	client := db.Initialize(&cfg)
	db.Migrate(client)
	initRouter(&cfg)
}
