package models

import (
	"time"

	"gorm.io/gorm"
)

// Category represents a category of prompts
type Category struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description string         `json:"description"`
	Color       string         `json:"color"`
	IconURL     string         `json:"icon_url"`
	PromptCount int64          `json:"prompt_count" gorm:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// CategoryCreate represents the data required to create a category
type CategoryCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
}

// CategoryUpdate represents the data required to update a category
type CategoryUpdate struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
}
