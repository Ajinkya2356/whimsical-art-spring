package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

var DB *gorm.DB

// InitializeDatabase initializes the GORM database connection
func InitializeDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	// Connect to the database with prepared statement settings
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		PrepareStmt:     false, // Disable prepared statements
		CreateBatchSize: 1000,
		QueryFields:     true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	fmt.Println("Database connection established")
	return DB, nil
}
