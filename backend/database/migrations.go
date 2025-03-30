package database

import (
	"backend/models"

	"gorm.io/gorm"
)

// RunMigrations runs all database migrations
func RunMigrations(db *gorm.DB) error {
	// Create tables with the new schema if they don't exist
	err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Prompt{},
		&models.Like{},
		&models.Favorite{},
	)
	if err != nil {
		return err
	}

	return nil
}
