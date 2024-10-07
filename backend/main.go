package main

import (
	"exchangeapp/backend/config"
	"exchangeapp/backend/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = "8000"
	}

	r.Run("0.0.0.0:" + port)
}
