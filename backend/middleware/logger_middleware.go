package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LoggerMiddleware logs incoming requests and their response times
func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Start timer
		start := time.Now()
		
		// Store the path
		path := c.Path()
		
		// Store the method
		method := c.Method()
		
		// Process request
		err := c.Next()
		
		// Calculate response time
		responseTime := time.Since(start)
		
		// Format log message
		logMessage := fmt.Sprintf(
			"[%s] %s - %s - %d - %s",
			time.Now().Format(time.RFC3339),
			method,
			path,
			c.Response().StatusCode(),
			responseTime,
		)
		
		// Print log message
		fmt.Println(logMessage)
		
		// Return error if any
		return err
	}
}
