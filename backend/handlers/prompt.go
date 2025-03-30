package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/models"
	"backend/repositories"
)

// PromptHandler handles prompt requests
type PromptHandler struct {
	promptRepo   repositories.PromptRepository
	categoryRepo repositories.CategoryRepository
	userRepo     repositories.UserRepository
}

// NewPromptHandler creates a new PromptHandler
func NewPromptHandler(promptRepo repositories.PromptRepository, categoryRepo repositories.CategoryRepository, userRepo repositories.UserRepository) *PromptHandler {
	return &PromptHandler{
		promptRepo:   promptRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

// GetPrompts gets all prompts with pagination and filtering
// @Summary Get all prompts
// @Description Retrieve a list of prompts with pagination and filtering options
// @Tags Prompts
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param pageSize query int false "Number of items per page"
// @Param categoryId query string false "Filter by category ID"
// @Param search query string false "Search term for filtering prompts"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string "Failed to get prompts"
// @Router /prompts [get]
func (h *PromptHandler) GetPrompts(c *gin.Context) {
	filter := models.PromptFilter{
		Page:     1,
		PageSize: 10,
	}

	prompts, total, err := h.promptRepo.GetPrompts(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prompts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"prompts": prompts,
		"total":   total,
	})
}

// GetPrompt gets a prompt by ID
// @Summary Get a specific prompt
// @Description Retrieve details of a specific prompt by ID
// @Tags Prompts
// @Produce json
// @Param id path string true "Prompt ID"
// @Success 200 {object} models.Prompt
// @Failure 400 {object} map[string]string "Invalid prompt ID"
// @Failure 404 {object} map[string]string "Prompt not found"
// @Router /prompts/{id} [get]
func (h *PromptHandler) GetPrompt(c *gin.Context) {
	id := c.Param("id")
	promptID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prompt ID"})
		return
	}

	prompt, err := h.promptRepo.GetPromptByID(promptID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Prompt not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prompt"})
		return
	}

	c.JSON(http.StatusOK, prompt)
}

// CreatePrompt creates a new prompt
// @Summary Create a new prompt
// @Description Create a new prompt with the provided information
// @Tags Prompts
// @Accept json
// @Produce json
// @Param prompt body models.PromptCreate true "Prompt to create"
// @Success 201 {object} models.Prompt
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 401 {object} map[string]string "Not authenticated"
// @Failure 500 {object} map[string]string "Failed to create prompt"
// @Security BearerAuth
// @Router /prompts [post]
func (h *PromptHandler) CreatePrompt(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var createData models.PromptCreate
	if err := c.ShouldBindJSON(&createData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	prompt, err := h.promptRepo.CreatePrompt(createData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create prompt"})
		return
	}

	c.JSON(http.StatusCreated, prompt)
}

// UpdatePrompt updates an existing prompt
// @Summary Update a prompt
// @Description Update an existing prompt's information
// @Tags Prompts
// @Accept json
// @Produce json
// @Param id path string true "Prompt ID"
// @Param prompt body models.PromptUpdate true "Updated prompt information"
// @Success 200 {object} models.Prompt
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 401 {object} map[string]string "Not authenticated"
// @Failure 403 {object} map[string]string "Not authorized to update this prompt"
// @Failure 404 {object} map[string]string "Prompt not found"
// @Failure 500 {object} map[string]string "Failed to update prompt"
// @Security BearerAuth
// @Router /prompts/{id} [put]
func (h *PromptHandler) UpdatePrompt(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return
	}

	id := c.Param("id")
	promptID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prompt ID"})
		return
	}

	prompt, err := h.promptRepo.GetPromptByID(promptID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Prompt not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prompt"})
		return
	}

	if prompt.UserID != currentUser.ID && !currentUser.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this prompt"})
		return
	}

	var updateData models.PromptUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedPrompt, err := h.promptRepo.UpdatePrompt(promptID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update prompt"})
		return
	}

	c.JSON(http.StatusOK, updatedPrompt)
}

// DeletePrompt deletes a prompt
// @Summary Delete a prompt
// @Description Delete an existing prompt
// @Tags Prompts
// @Param id path string true "Prompt ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid prompt ID"
// @Failure 401 {object} map[string]string "Not authenticated"
// @Failure 403 {object} map[string]string "Not authorized to delete this prompt"
// @Failure 404 {object} map[string]string "Prompt not found"
// @Failure 500 {object} map[string]string "Failed to delete prompt"
// @Security BearerAuth
// @Router /prompts/{id} [delete]
func (h *PromptHandler) DeletePrompt(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return
	}

	id := c.Param("id")
	promptID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prompt ID"})
		return
	}

	prompt, err := h.promptRepo.GetPromptByID(promptID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Prompt not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prompt"})
		return
	}

	if prompt.UserID != currentUser.ID && !currentUser.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this prompt"})
		return
	}

	if err := h.promptRepo.DeletePrompt(promptID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete prompt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prompt deleted successfully"})
}

// LikePrompt toggles a like on a prompt by the current user
// @Summary Toggle like on a prompt
// @Description Like or unlike a prompt for the authenticated user
// @Tags Prompts
// @Produce json
// @Param id path string true "Prompt ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid prompt ID"
// @Failure 401 {object} map[string]string "Not authenticated"
// @Failure 500 {object} map[string]string "Failed to process like"
// @Security BearerAuth
// @Router /prompts/{id}/like [post]
func (h *PromptHandler) LikePrompt(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return
	}

	id := c.Param("id")
	promptID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prompt ID"})
		return
	}

	prompt, err := h.promptRepo.GetPromptByID(promptID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Prompt not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prompt"})
		return
	}

	like := &models.Like{
		UserID:   currentUser.ID,
		PromptID: prompt.ID,
	}

	if err := h.promptRepo.CreateLike(like); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like prompt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prompt liked successfully"})
}

// GetPromptsByCategory gets all prompts for a specific category
// @Summary Get prompts in a category
// @Description Retrieve all prompts belonging to a specific category
// @Tags Categories,Prompts
// @Produce json
// @Param id path string true "Category ID"
// @Param page query int false "Page number for pagination"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid category ID"
// @Failure 500 {object} map[string]string "Failed to get prompts"
// @Router /categories/{id}/prompts [get]
func (h *PromptHandler) GetPromptsByCategory(c *gin.Context) {
	categoryID := c.Param("id")
	prompts, err := h.promptRepo.GetPromptsByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prompts"})
		return
	}

	c.JSON(http.StatusOK, prompts)
}
