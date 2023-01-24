package main

import (
	"api/server"
)

func main() {
	s := server.InitializeServer()

	s.Run()
}
