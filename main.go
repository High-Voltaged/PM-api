package main

import (
	"api/server"
)

func main() {
	a := server.InitializeApp()

	a.Run()
}
