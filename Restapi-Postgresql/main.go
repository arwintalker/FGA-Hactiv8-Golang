package main

import (
	"Restapi-Postgresql/config"
	"Restapi-Postgresql/routers"
)

func main() {
	var PORT = ":8080"

	config.ConnectDB()

	routers.StartServer().Run(PORT)
}
