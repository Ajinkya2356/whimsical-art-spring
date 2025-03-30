package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/models"
)

// CategoryRepository is the interface for category operations
type CategoryRepository interface {
	GetCategories() ([]models.Category, error)
	GetCategoryByID(id uuid.UUID) (*models.Category, error)
	GetCategoryByName(name string) (*models.Category, error)
	CreateCategory(create models.CategoryCreate) (*models.Category, error)
	UpdateCategory(id uuid.UUID, update models.CategoryUpdate) (*models.Category, error)
	DeleteCategory(id uuid.UUID) error
}

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) CategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (r *GormCategoryRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *GormCategoryRepository) GetCategoryByID(id uuid.UUID) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *GormCategoryRepository) GetCategoryByName(name string) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *GormCategoryRepository) CreateCategory(create models.CategoryCreate) (*models.Category, error) {
	category := &models.Category{
		Name:        create.Name,
		Description: create.Description,
		Color:       create.Color,
		IconURL:     create.IconURL,
	}

	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *GormCategoryRepository) UpdateCategory(id uuid.UUID, update models.CategoryUpdate) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if update.Name != nil {
		updates["name"] = *update.Name
	}
	if update.Description != nil {
		updates["description"] = *update.Description
	}
	if update.Color != nil {
		updates["color"] = *update.Color
	}
	if update.IconURL != nil {
		updates["icon_url"] = *update.IconURL
	}
	updates["updated_at"] = time.Now()

	if err := r.db.Model(&category).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *GormCategoryRepository) DeleteCategory(id uuid.UUID) error {
	return r.db.Delete(&models.Category{}, "id = ?", id).Error
}
