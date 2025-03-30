package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggingMiddleware handles request logging
type LoggingMiddleware struct {
	Logger *logrus.Logger
}

// NewLoggingMiddleware creates a new LoggingMiddleware
func NewLoggingMiddleware(logger *logrus.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		Logger: logger,
	}
}

// LogRequest logs request details
func (l *LoggingMiddleware) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()

		// Create log entry
		entry := l.Logger.WithFields(logrus.Fields{
			"method":     method,
			"path":       path,
			"status":     status,
			"duration":   duration,
			"ip":         ip,
			"user_agent": c.Request.UserAgent(),
		})

		// Log based on status code
		if status >= 500 {
			entry.Error("Server error")
		} else if status >= 400 {
			entry.Warn("Client error")
		} else {
			entry.Info("Request processed")
		}
	}
}
