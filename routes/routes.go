package routes

import (
	"authentication/controllers"
	"authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup())
	router.POST("/login", controllers.Login())

	protected := router.Group("/")
	protected.Use(middlewares.Authenticate())
	{
		protected.GET("/users", controllers.GetUsers())
		protected.GET("/users/:id", controllers.GetUser())
	}
}
