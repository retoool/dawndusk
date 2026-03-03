package handlers

import (
	"net/http"
	"strconv"

	"github.com/dawndusk/backend/internal/api/middlewares"
	"github.com/dawndusk/backend/internal/domain/services"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DecorationHandler struct {
	decorationService services.DecorationService
}

func NewDecorationHandler(decorationService services.DecorationService) *DecorationHandler {
	return &DecorationHandler{
		decorationService: decorationService,
	}
}

// ListDecorations godoc
// @Summary List all decorations
// @Description Get list of all available decorations
// @Tags decorations
// @Produce json
// @Param limit query int false "Limit" default(50)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.DecorationResponse
// @Router /api/v1/decorations [get]
func (h *DecorationHandler) ListDecorations(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	responses, err := h.decorationService.ListDecorations(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// GetUserDecorations godoc
// @Summary Get user's decorations
// @Description Get all decorations unlocked by the user
// @Tags decorations
// @Produce json
// @Success 200 {array} dto.UserDecorationResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/decorations/my [get]
func (h *DecorationHandler) GetUserDecorations(c *gin.Context) {
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

	responses, err := h.decorationService.GetUserDecorations(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// UnlockDecoration godoc
// @Summary Unlock a decoration
// @Description Unlock a decoration for the user
// @Tags decorations
// @Accept json
// @Produce json
// @Param request body dto.UnlockDecorationRequest true "Unlock decoration request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/decorations/unlock [post]
func (h *DecorationHandler) UnlockDecoration(c *gin.Context) {
	var req dto.UnlockDecorationRequest
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

	decorationID, err := uuid.Parse(req.DecorationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid decoration ID"})
		return
	}

	if err := h.decorationService.UnlockDecoration(userID, decorationID); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Decoration unlocked successfully"})
}

// EquipDecoration godoc
// @Summary Equip a decoration
// @Description Equip a decoration for the user's pet
// @Tags decorations
// @Param id path string true "Decoration ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/decorations/{id}/equip [post]
func (h *DecorationHandler) EquipDecoration(c *gin.Context) {
	decorationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid decoration ID"})
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

	if err := h.decorationService.EquipDecoration(userID, decorationID); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Decoration equipped successfully"})
}
