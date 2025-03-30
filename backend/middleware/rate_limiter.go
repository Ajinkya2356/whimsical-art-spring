package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/models"
)

// RateLimiter handles rate limiting
type RateLimiter struct {
	DB            *gorm.DB
	Logger        *logrus.Logger
	globalLimiter map[string][]time.Time
	promptLimiter map[string][]time.Time
	globalMutex   sync.RWMutex
	promptMutex   sync.RWMutex
	cleanupTicker *time.Ticker
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(db *gorm.DB, logger *logrus.Logger) *RateLimiter {
	limiter := &RateLimiter{
		DB:            db,
		Logger:        logger,
		globalLimiter: make(map[string][]time.Time),
		promptLimiter: make(map[string][]time.Time),
		cleanupTicker: time.NewTicker(time.Minute),
	}

	// Start cleanup goroutine
	go limiter.cleanupLoop()

	return limiter
}

// LimitGlobalRequests limits the number of requests per IP
func (r *RateLimiter) LimitGlobalRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get client IP
		clientIP := c.ClientIP()

		// Check rate limit
		r.globalMutex.Lock()
		now := time.Now()

		// Remove old requests
		requests := r.globalLimiter[clientIP]
		validRequests := requests[:0]
		for _, t := range requests {
			if now.Sub(t) < time.Minute {
				validRequests = append(validRequests, t)
			}
		}

		// Check if limit exceeded
		if len(validRequests) >= 100 {
			r.globalMutex.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		// Add new request
		validRequests = append(validRequests, now)
		r.globalLimiter[clientIP] = validRequests
		r.globalMutex.Unlock()

		c.Next()
	}
}

// LimitPromptCreation limits the number of prompts created per user
func (r *RateLimiter) LimitPromptCreation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user from context
		user, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Get user ID
		userID := user.(*models.User).ID

		// Check rate limit
		r.promptMutex.Lock()
		now := time.Now()

		// Remove old requests
		requests := r.promptLimiter[userID]
		validRequests := requests[:0]
		for _, t := range requests {
			if now.Sub(t) < time.Hour {
				validRequests = append(validRequests, t)
			}
		}

		// Check if limit exceeded
		if len(validRequests) >= 10 {
			r.promptMutex.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many prompt creations"})
			c.Abort()
			return
		}

		// Add new request
		validRequests = append(validRequests, now)
		r.promptLimiter[userID] = validRequests
		r.promptMutex.Unlock()

		c.Next()
	}
}

// cleanup removes old requests
func (r *RateLimiter) cleanup() {
	now := time.Now()

	// Clean up global limiter
	r.globalMutex.Lock()
	for ip, requests := range r.globalLimiter {
		validRequests := requests[:0]
		for _, t := range requests {
			if now.Sub(t) < time.Minute {
				validRequests = append(validRequests, t)
			}
		}
		if len(validRequests) == 0 {
			delete(r.globalLimiter, ip)
		} else {
			r.globalLimiter[ip] = validRequests
		}
	}
	r.globalMutex.Unlock()

	// Clean up prompt limiter
	r.promptMutex.Lock()
	for userID, requests := range r.promptLimiter {
		validRequests := requests[:0]
		for _, t := range requests {
			if now.Sub(t) < time.Hour {
				validRequests = append(validRequests, t)
			}
		}
		if len(validRequests) == 0 {
			delete(r.promptLimiter, userID)
		} else {
			r.promptLimiter[userID] = validRequests
		}
	}
	r.promptMutex.Unlock()
}

// cleanupLoop runs the cleanup periodically
func (r *RateLimiter) cleanupLoop() {
	for range r.cleanupTicker.C {
		r.cleanup()
	}
}
