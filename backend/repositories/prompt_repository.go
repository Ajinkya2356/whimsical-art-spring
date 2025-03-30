package repositories

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/models"
)

// PromptRepository is the interface for prompt operations
type PromptRepository interface {
	GetPrompts(filter models.PromptFilter) ([]models.PromptResponse, int, error)
	GetPromptByID(id uuid.UUID) (*models.PromptResponse, error)
	CreatePrompt(create models.PromptCreate) (*models.Prompt, error)
	UpdatePrompt(id uuid.UUID, update models.PromptUpdate) (*models.Prompt, error)
	DeletePrompt(id uuid.UUID) error
	AddPromptView(promptID uuid.UUID, userID *uuid.UUID, ip, userAgent string) error
	AddPromptFavorite(promptID, userID uuid.UUID) error
	RemovePromptFavorite(promptID, userID uuid.UUID) error
	GetUserFavorites(userID uuid.UUID) ([]models.PromptResponse, error)
	IsPromptFavoritedByUser(promptID, userID uuid.UUID) (bool, error)
	ToggleLike(promptID, userID uuid.UUID) (bool, error)
	CreateLike(like *models.Like) error
	GetPromptsByCategory(categoryID string) ([]models.Prompt, error)
}

type GormPromptRepository struct {
	db *gorm.DB
}

func NewGormPromptRepository(db *gorm.DB) PromptRepository {
	return &GormPromptRepository{db: db}
}

func (r *GormPromptRepository) GetPrompts(filter models.PromptFilter) ([]models.PromptResponse, int, error) {
	var prompts []models.PromptResponse
	var total int64

	query := r.db.Model(&models.Prompt{})

	// Apply filters
	if filter.Search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if filter.CategoryID != "" {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.PageSize
	query = query.Offset(offset).Limit(filter.PageSize)

	// Apply sorting
	if filter.SortBy != "" {
		order := "DESC"
		if !filter.SortDesc {
			order = "ASC"
		}
		query = query.Order(filter.SortBy + " " + order)
	} else {
		query = query.Order("created_at DESC")
	}

	// Execute query
	if err := query.Find(&prompts).Error; err != nil {
		return nil, 0, err
	}

	return prompts, int(total), nil
}

func (r *GormPromptRepository) GetPromptByID(id uuid.UUID) (*models.PromptResponse, error) {
	var prompt models.PromptResponse
	if err := r.db.First(&prompt, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &prompt, nil
}

func (r *GormPromptRepository) CreatePrompt(create models.PromptCreate) (*models.Prompt, error) {
	prompt := &models.Prompt{
		Title:       create.Title,
		Description: create.Description,
		Content:     create.Content,
		CategoryID:  create.CategoryID,
		ImageURL:    create.ImageURL,
	}

	if err := r.db.Create(prompt).Error; err != nil {
		return nil, err
	}
	return prompt, nil
}

func (r *GormPromptRepository) UpdatePrompt(id uuid.UUID, update models.PromptUpdate) (*models.Prompt, error) {
	var prompt models.Prompt
	if err := r.db.First(&prompt, "id = ?", id).Error; err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if update.Title != nil {
		updates["title"] = *update.Title
	}
	if update.Description != nil {
		updates["description"] = *update.Description
	}
	if update.Content != nil {
		updates["content"] = *update.Content
	}
	if update.CategoryID != nil {
		updates["category_id"] = *update.CategoryID
	}
	if update.ImageURL != nil {
		updates["image_url"] = *update.ImageURL
	}
	updates["updated_at"] = time.Now()

	if err := r.db.Model(&prompt).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &prompt, nil
}

func (r *GormPromptRepository) DeletePrompt(id uuid.UUID) error {
	return r.db.Delete(&models.Prompt{}, "id = ?", id).Error
}

func (r *GormPromptRepository) CreateLike(like *models.Like) error {
	return r.db.Create(like).Error
}

func (r *GormPromptRepository) GetPromptsByCategory(categoryID string) ([]models.Prompt, error) {
	var prompts []models.Prompt
	if err := r.db.Where("category_id = ?", categoryID).Find(&prompts).Error; err != nil {
		return nil, err
	}
	return prompts, nil
}

// AddPromptView adds a view to a prompt
func (r *GormPromptRepository) AddPromptView(promptID uuid.UUID, userID *uuid.UUID, ip, userAgent string) error {
	// Check if prompt exists
	_, err := r.GetPromptByID(promptID)
	if err != nil {
		return err
	}

	// Insert view
	if err := r.db.Exec(`
		INSERT INTO views (prompt_id, user_id, ip, user_agent)
		VALUES (?, ?, ?, ?)
	`, promptID, userID, ip, userAgent).Error; err != nil {
		return err
	}

	return nil
}

// AddPromptFavorite adds a favorite to a prompt
func (r *GormPromptRepository) AddPromptFavorite(promptID, userID uuid.UUID) error {
	// Check if prompt exists
	_, err := r.GetPromptByID(promptID)
	if err != nil {
		return err
	}

	// Check if user exists
	var userExists bool
	if err := r.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", userID).Scan(&userExists).Error; err != nil {
		return err
	}
	if !userExists {
		return errors.New("user not found")
	}

	// Check if already favorited
	var favoriteExists bool
	if err := r.db.Raw(`
		SELECT EXISTS(SELECT 1 FROM favorites WHERE prompt_id = ? AND user_id = ?)
	`, promptID, userID).Scan(&favoriteExists).Error; err != nil {
		return err
	}
	if favoriteExists {
		return errors.New("prompt already favorited")
	}

	// Insert favorite
	if err := r.db.Exec(`
		INSERT INTO favorites (prompt_id, user_id)
		VALUES (?, ?)
	`, promptID, userID).Error; err != nil {
		return err
	}

	return nil
}

// RemovePromptFavorite removes a favorite from a prompt
func (r *GormPromptRepository) RemovePromptFavorite(promptID, userID uuid.UUID) error {
	// Check if favorite exists
	var favoriteExists bool
	if err := r.db.Raw(`
		SELECT EXISTS(SELECT 1 FROM favorites WHERE prompt_id = ? AND user_id = ?)
	`, promptID, userID).Scan(&favoriteExists).Error; err != nil {
		return err
	}
	if !favoriteExists {
		return errors.New("favorite not found")
	}

	// Delete favorite
	if err := r.db.Exec(`
		DELETE FROM favorites WHERE prompt_id = ? AND user_id = ?
	`, promptID, userID).Error; err != nil {
		return err
	}

	return nil
}

// GetUserFavorites gets prompts favorited by a user
func (r *GormPromptRepository) GetUserFavorites(userID uuid.UUID) ([]models.PromptResponse, error) {
	var prompts []models.PromptResponse
	if err := r.db.Model(&models.Prompt{}).
		Joins("JOIN favorites f ON prompts.id = f.prompt_id").
		Where("f.user_id = ?", userID).
		Select("prompts.*").
		Find(&prompts).Error; err != nil {
		return nil, err
	}
	return prompts, nil
}

// IsPromptFavoritedByUser checks if a prompt is favorited by a user
func (r *GormPromptRepository) IsPromptFavoritedByUser(promptID, userID uuid.UUID) (bool, error) {
	var favorited bool
	if err := r.db.Raw(`
		SELECT EXISTS(SELECT 1 FROM favorites WHERE prompt_id = ? AND user_id = ?)
	`, promptID, userID).Scan(&favorited).Error; err != nil {
		return false, err
	}
	return favorited, nil
}

// ToggleLike toggles a like (favorite) on a prompt for a user
// Returns the new state (true = liked, false = unliked)
func (r *GormPromptRepository) ToggleLike(promptID, userID uuid.UUID) (bool, error) {
	// Check if prompt exists
	_, err := r.GetPromptByID(promptID)
	if err != nil {
		return false, err
	}

	// Check current state
	currentlyFavorited, err := r.IsPromptFavoritedByUser(promptID, userID)
	if err != nil {
		return false, err
	}

	if currentlyFavorited {
		// Remove favorite
		if err := r.db.Exec(`
			DELETE FROM favorites WHERE prompt_id = ? AND user_id = ?
		`, promptID, userID).Error; err != nil {
			return false, err
		}

		// Decrement favorite count
		if err := r.db.Exec(`
			UPDATE prompts
			SET favorite_count = GREATEST(0, favorite_count - 1)
			WHERE id = ?
		`, promptID).Error; err != nil {
			return false, err
		}
	} else {
		// Add favorite
		if err := r.db.Exec(`
			INSERT INTO favorites (prompt_id, user_id)
			VALUES (?, ?)
		`, promptID, userID).Error; err != nil {
			return false, err
		}

		// Increment favorite count
		if err := r.db.Exec(`
			UPDATE prompts
			SET favorite_count = favorite_count + 1
			WHERE id = ?
		`, promptID).Error; err != nil {
			return false, err
		}
	}
	// Return the new state (opposite of current state)
	return !currentlyFavorited, nil
}
