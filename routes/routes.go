package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigosuco/go-url-shortener/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Application up and running!",
		})
	})

	router.POST("/url", controllers.CreateShortUrl)
}