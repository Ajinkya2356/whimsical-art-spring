package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Prompt represents an AI prompt
type Prompt struct {
	ID            string         `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"not null"`
	Description   string         `json:"description"`
	Content       string         `json:"content"`
	CategoryID    string         `json:"category_id" gorm:"not null"`
	UserID        string         `json:"user_id" gorm:"not null"`
	ViewCount     int64          `json:"view_count" gorm:"default:0"`
	FavoriteCount int64          `json:"favorite_count" gorm:"default:0"`
	ImageURL      string         `json:"image_url"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// PromptCreate represents the data required to create a prompt
type PromptCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content" binding:"required"`
	CategoryID  string `json:"category_id" binding:"required"`
	ImageURL    string `json:"image_url,omitempty"`
}

// PromptUpdate represents the data required to update a prompt
type PromptUpdate struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Content     *string `json:"content,omitempty"`
	CategoryID  *string `json:"category_id,omitempty"`
	ImageURL    *string `json:"image_url,omitempty"`
}

// PromptResponse represents a prompt with joined data
type PromptResponse struct {
	Prompt
	Username    string `json:"username"`
	Category    string `json:"category"`
	IsFavorited bool   `json:"is_favorited,omitempty"`
}

// PromptFilter represents filters for searching prompts
type PromptFilter struct {
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"page_size" form:"page_size"`
	Search     string `json:"search" form:"search"`
	CategoryID string `json:"category_id" form:"category_id"`
	UserID     string `json:"user_id" form:"user_id"`
	SortBy     string `json:"sort_by" form:"sort_by"`
	SortDesc   bool   `json:"sort_desc" form:"sort_desc"`
}

// PromptFavorite represents a user favoriting a prompt
type PromptFavorite struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	PromptID  uuid.UUID `json:"promptId"`
	CreatedAt time.Time `json:"createdAt"`
}

// PromptView represents a view of a prompt
type PromptView struct {
	ID        uuid.UUID  `json:"id"`
	UserID    *uuid.UUID `json:"userId,omitempty"` // Optional, for logged in users
	PromptID  uuid.UUID  `json:"promptId"`
	CreatedAt time.Time  `json:"createdAt"`
	IP        string     `json:"ip"`
	UserAgent string     `json:"userAgent"`
}
