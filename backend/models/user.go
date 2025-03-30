package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user of the system
type User struct {
	ID        string         `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	IsAdmin   bool           `json:"is_admin" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	LastLogin time.Time      `json:"last_login"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// UserCreate represents the data required to create a user
type UserCreate struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserUpdate represents the data required to update a user
type UserUpdate struct {
	Email     *string    `json:"email,omitempty" binding:"omitempty,email"`
	Username  *string    `json:"username,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// UserStats represents user statistics
type UserStats struct {
	TotalUsers         int64 `json:"total_users"`
	TotalPrompts       int64 `json:"total_prompts"`
	TotalFavorites     int64 `json:"total_favorites"`
	TotalViews         int64 `json:"total_views"`
	NewUsersLast24h    int64 `json:"new_users_last_24h"`
	NewUsersLast7d     int64 `json:"new_users_last_7d"`
	NewUsersLast30d    int64 `json:"new_users_last_30d"`
	ActiveUsersLast24h int64 `json:"active_users_last_24h"`
	ActiveUsersLast7d  int64 `json:"active_users_last_7d"`
	ActiveUsersLast30d int64 `json:"active_users_last_30d"`
	TotalLikesReceived int64 `json:"total_likes_received"`
	TotalLikesGiven    int64 `json:"total_likes_given"`
	RecentActivity     int64 `json:"recent_activity"`
}

// SystemStats represents system-wide statistics
type SystemStats struct {
	TotalUsers      int64 `json:"total_users"`
	TotalPrompts    int64 `json:"total_prompts"`
	TotalTags       int64 `json:"total_tags"`
	TotalCategories int64 `json:"total_categories"`
	NewUsers7Days   int64 `json:"new_users_7days"`
	NewPrompts7Days int64 `json:"new_prompts_7days"`
	TotalViews7Days int64 `json:"total_views_7days"`
	TotalLikes      int64 `json:"total_likes"`
	RecentActivity  int64 `json:"recent_activity"`
}

// UserResponse represents a user with joined data and computed fields
type UserResponse struct {
	User
	PromptCount   int64 `gorm:"-" json:"promptCount"`   // Ignored by GORM
	FavoriteCount int64 `gorm:"-" json:"favoriteCount"` // Ignored by GORM
}

// Favorite represents a user's favorite prompt
type Favorite struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	UserID    string    `gorm:"type:uuid;not null"`
	PromptID  string    `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Like represents a user's like on a prompt
type Like struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	UserID    string         `json:"user_id" gorm:"not null"`
	PromptID  string         `json:"prompt_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
