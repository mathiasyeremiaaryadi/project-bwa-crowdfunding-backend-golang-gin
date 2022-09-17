package main

import (
	"fmt"
	"service-campaign-startup/app"
	"service-campaign-startup/config"
)

func main() {
	config.InitConfig()

	db := config.ConnectMySQL()

	router := app.InitRoute(db)

	appHost := fmt.Sprintf("localhost:%s", config.CONFIG["PORT"])
	router.Run(appHost)
}
