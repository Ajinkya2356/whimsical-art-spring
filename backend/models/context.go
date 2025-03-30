package models

// ContextKey is a type used for context keys to avoid collisions
type ContextKey string

// Context keys used in the application
const (
	// ContextKeyUserID is the key for storing the user ID in the context
	ContextKeyUserID ContextKey = "userID"
	
	// ContextKeyUser is the key for storing the user object in the context
	ContextKeyUser ContextKey = "user"
)