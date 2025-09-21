package main

import (
	"fmt"
	"service-campaign-startup/app"
	"service-campaign-startup/config"

	"github.com/spf13/viper"
)

func main() {
	dependencies := config.NewDepenencies()

	applicationServer := fmt.Sprintf("localhost:%s", viper.GetString("APP_PORT"))

	router := app.NewRoute(dependencies)
	router.Run(applicationServer)
}
