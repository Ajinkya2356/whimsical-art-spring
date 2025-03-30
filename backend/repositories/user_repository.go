package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"backend/models"
)

// UserRepository is the interface for user operations
type UserRepository interface {
	GetUsers() ([]models.UserResponse, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	GetUserStats(userID string) (*models.UserStats, error)
	GetSystemStats() (*models.SystemStats, error)
	GenerateToken(userID string) (string, error)
	VerifyToken(token string) (string, error)
	UpdateLastLogin(userID string) error
}

// GormUserRepository implements UserRepository using GORM
type GormUserRepository struct {
	DB *gorm.DB
}

// NewGormUserRepository creates a new GormUserRepository
func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{
		DB: db,
	}
}

// GetUsers retrieves all users with their prompt and favorite counts
func (r *GormUserRepository) GetUsers() ([]models.UserResponse, error) {
	var users []models.UserResponse
	err := r.DB.Model(&models.User{}).
		Select(`
			users.id, users.email, users.username,
			users.created_at, users.updated_at, users.last_login, users.is_admin,
			COUNT(DISTINCT prompts.id) AS prompt_count,
			COUNT(DISTINCT favorites.prompt_id) AS favorite_count
		`).
		Joins("LEFT JOIN prompts ON users.id = prompts.user_id").
		Joins("LEFT JOIN favorites ON users.id = favorites.user_id").
		Group("users.id").
		Order("users.created_at DESC").
		Scan(&users).Error
	return users, err
}

// GetUserByID retrieves a user by their ID
func (r *GormUserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email
func (r *GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (r *GormUserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// UpdateUser updates an existing user
func (r *GormUserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

// DeleteUser deletes a user and their related data
func (r *GormUserRepository) DeleteUser(id string) error {
	return r.DB.Where("id = ?", id).Delete(&models.User{}).Error
}

// GetUserStats retrieves user statistics
func (r *GormUserRepository) GetUserStats(userID string) (*models.UserStats, error) {
	var stats models.UserStats

	// Get total prompts
	if err := r.DB.Model(&models.Prompt{}).Where("user_id = ?", userID).Count(&stats.TotalPrompts).Error; err != nil {
		return nil, err
	}

	// Get total likes received
	if err := r.DB.Model(&models.Like{}).Where("prompt_id IN (SELECT id FROM prompts WHERE user_id = ?)", userID).Count(&stats.TotalLikesReceived).Error; err != nil {
		return nil, err
	}

	// Get total likes given
	if err := r.DB.Model(&models.Like{}).Where("user_id = ?", userID).Count(&stats.TotalLikesGiven).Error; err != nil {
		return nil, err
	}

	// Get recent activity (last 7 days)
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	if err := r.DB.Model(&models.Prompt{}).Where("user_id = ? AND created_at >= ?", userID, sevenDaysAgo).Count(&stats.RecentActivity).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetSystemStats retrieves system statistics
func (r *GormUserRepository) GetSystemStats() (*models.SystemStats, error) {
	var stats models.SystemStats

	// Get total users
	if err := r.DB.Model(&models.User{}).Count(&stats.TotalUsers).Error; err != nil {
		return nil, err
	}

	// Get total prompts
	if err := r.DB.Model(&models.Prompt{}).Count(&stats.TotalPrompts).Error; err != nil {
		return nil, err
	}

	// Get total categories
	if err := r.DB.Model(&models.Category{}).Count(&stats.TotalCategories).Error; err != nil {
		return nil, err
	}

	// Get total likes
	if err := r.DB.Model(&models.Like{}).Count(&stats.TotalLikes).Error; err != nil {
		return nil, err
	}

	// Get recent activity (last 7 days)
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	if err := r.DB.Model(&models.Prompt{}).Where("created_at >= ?", sevenDaysAgo).Count(&stats.RecentActivity).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

// GenerateToken generates a JWT token for a user
func (r *GormUserRepository) GenerateToken(userID string) (string, error) {
	// TODO: Implement JWT token generation
	return "", errors.New("not implemented")
}

// VerifyToken verifies a JWT token and returns the user ID
func (r *GormUserRepository) VerifyToken(token string) (string, error) {
	// TODO: Implement JWT token verification
	return "", errors.New("not implemented")
}

// UpdateLastLogin updates the last login time for a user
func (r *GormUserRepository) UpdateLastLogin(userID string) error {
	return r.DB.Model(&models.User{}).Where("id = ?", userID).Update("last_login", time.Now()).Error
}
