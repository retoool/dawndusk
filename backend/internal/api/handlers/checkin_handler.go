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

type CheckInHandler struct {
	checkInService services.CheckInService
}

func NewCheckInHandler(checkInService services.CheckInService) *CheckInHandler {
	return &CheckInHandler{
		checkInService: checkInService,
	}
}

// Create godoc
// @Summary Create a new check-in
// @Description Create a new wake or sleep check-in
// @Tags check-ins
// @Accept json
// @Produce json
// @Param request body dto.CreateCheckInRequest true "Check-in request"
// @Success 201 {object} dto.CheckInResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/check-ins [post]
func (h *CheckInHandler) Create(c *gin.Context) {
	var req dto.CreateCheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from auth middleware
	userIDStr, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse user ID to UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	response, err := h.checkInService.CreateCheckIn(userID, &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetList godoc
// @Summary Get check-in history
// @Description Get paginated list of user's check-ins
// @Tags check-ins
// @Produce json
// @Param limit query int false "Limit" default(20)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.CheckInResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/check-ins [get]
func (h *CheckInHandler) GetList(c *gin.Context) {
	// Get user ID from auth middleware
	userIDStr, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse user ID to UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse query parameters
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	response, err := h.checkInService.GetCheckIns(userID, limit, offset)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetToday godoc
// @Summary Get today's check-ins
// @Description Get today's wake and sleep check-ins
// @Tags check-ins
// @Produce json
// @Success 200 {object} dto.TodayCheckInsResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/check-ins/today [get]
func (h *CheckInHandler) GetToday(c *gin.Context) {
	// Get user ID from auth middleware
	userIDStr, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse user ID to UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	response, err := h.checkInService.GetTodayCheckIns(userID)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetStats godoc
// @Summary Get check-in statistics
// @Description Get user's check-in statistics including streaks
// @Tags check-ins
// @Produce json
// @Param days query int false "Number of days" default(30)
// @Success 200 {object} dto.CheckInStatsResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/check-ins/stats [get]
func (h *CheckInHandler) GetStats(c *gin.Context) {
	// Get user ID from auth middleware
	userIDStr, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse user ID to UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse query parameter
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))

	response, err := h.checkInService.GetCheckInStats(userID, days)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}
