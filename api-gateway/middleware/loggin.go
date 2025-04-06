package middleware

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// loggin middleware
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("x-request-id", requestID)

		start := time.Now()
		c.Next()

		latency := time.Since(start)

		log.Printf("requestID: %s, method: %s, path: %s, latency: %v", requestID, c.Request.Method, c.Request.URL.Path, latency)
	}
}