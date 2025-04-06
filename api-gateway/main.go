package main

import (
	"api-gateway/handlers"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Add middleware logging
	r.Use(middleware.Logging())

	// define router to microservices
	r.Any("/auth/*path", handlers.ProxyHandler("http://localhost:8081"))
	r.Any("/user/*path", handlers.ProxyHandler("http://localhost:8082"))
	r.Any("/chat/*path", handlers.ProxyHandler("http://localhost:8083"))

	// router check connection
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}