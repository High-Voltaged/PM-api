package main

import (
	db "api/database"
	"api/routes"
	"api/utils"
)

func main() {
	var cfg = utils.ReadEnv()

	db.Connect(&cfg)

	// routes.InitializeRouter(&cfg)
}
