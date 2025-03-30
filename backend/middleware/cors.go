package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles CORS
type CORSMiddleware struct{}

// NewCORSMiddleware creates a new CORSMiddleware
func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

// CORS handles CORS headers
func (m *CORSMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24 hours

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
