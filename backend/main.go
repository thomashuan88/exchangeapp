package main

import (
	"exchangeapp/backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	// simple any
	type Info struct {
		Message string
	}
	InfoTest := Info{
		Message: "Hello World!",
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, InfoTest)
	})

	r.Run("0.0.0.0:" + config.AppConfig.App.Port)
}
