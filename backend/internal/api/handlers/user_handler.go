package handlers

import (
	"net/http"

	"github.com/dawndusk/backend/internal/api/middlewares"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userRepo repositories.UserRepository
}

func NewUserHandler(userRepo repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get current user's profile information
// @Tags users
// @Produce json
// @Success 200 {object} dto.UserProfileResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/users/me [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userIDStr, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response := &dto.UserProfileResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		AvatarURL:   user.AvatarURL,
		Timezone:    user.Timezone,
		IsVerified:  user.IsVerified,
		CreatedAt:   user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update current user's profile information
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UpdateUserProfileRequest true "Update profile request"
// @Success 200 {object} dto.UserProfileResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/users/me [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDStr, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update fields if provided
	if req.Username != nil {
		// Check if username is already taken
		existingUser, err := h.userRepo.FindByUsername(*req.Username)
		if err == nil && existingUser.ID != user.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
			return
		}
		user.Username = *req.Username
	}
	if req.PhoneNumber != nil {
		user.PhoneNumber = req.PhoneNumber
	}
	if req.AvatarURL != nil {
		user.AvatarURL = req.AvatarURL
	}
	if req.Timezone != nil {
		user.Timezone = *req.Timezone
	}

	if err := h.userRepo.Update(user); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	response := &dto.UserProfileResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		AvatarURL:   user.AvatarURL,
		Timezone:    user.Timezone,
		IsVerified:  user.IsVerified,
		CreatedAt:   user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}
