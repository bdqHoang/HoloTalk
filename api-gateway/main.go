package main

import (
	"api-gateway/config"
	"api-gateway/middleware"
	"api-gateway/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := router.RouterConfig()

	env := config.LoadEnv()

	// Add middleware logging
	r.Use(middleware.Logging())

	// router check connection
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Println("Server running on port " + env.PORT)
	r.Run(":" + env.PORT)
}