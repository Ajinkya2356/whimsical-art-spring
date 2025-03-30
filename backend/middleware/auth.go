package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"backend/models"
)

// AuthMiddleware handles authentication
type AuthMiddleware struct {
	db *gorm.DB
}

// NewAuthMiddleware creates a new AuthMiddleware
func NewAuthMiddleware(db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{
		db: db,
	}
}

// RequireAuth ensures that the request is authenticated
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify token
		claims, err := m.verifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Get user from database
		var user models.User
		if err := m.db.First(&user, "id = ?", claims.Subject).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", &user)
		c.Next()
	}
}

// Optional ensures that the request is optionally authenticated
func (m *AuthMiddleware) Optional() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify token
		claims, err := m.verifyToken(token)
		if err != nil {
			c.Next()
			return
		}

		// Get user from database
		var user models.User
		if err := m.db.First(&user, "id = ?", claims.Subject).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.Next()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", &user)
		c.Next()
	}
}

func (m *AuthMiddleware) verifyToken(tokenString string) (*jwt.RegisteredClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		// Return secret key
		return []byte("your-secret-key"), nil // In production, use environment variable
	})

	if err != nil {
		return nil, err
	}

	// Check if token is valid
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
