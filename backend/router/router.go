package router

import (
	"exchangeapp/backend/controllers"
	"exchangeapp/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchange_rate", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchange_rate", controllers.CreateExchangeRate)
		api.POST("/article", controllers.CreateArticle)
		api.GET("/article", controllers.GetArticle)
		api.GET("/article/:id", controllers.GetArticleById)
		api.POST("/article/:id/like", controllers.LikeArticle)
		api.GET("/article/:id/likes", controllers.GetArtileLikes)
	}

	return r
}
