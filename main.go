package main

import (
	"ApiBooth/config"
	"ApiBooth/router"
	"log"
)

func main() {
	config.ConnectDB()

	r := router.SetupRouter()
	log.Fatal(r.Run(":8888"))
}
