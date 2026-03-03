package handlers

import (
	"net/http"

	"github.com/dawndusk/backend/internal/api/middlewares"
	"github.com/dawndusk/backend/internal/domain/services"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SleepScheduleHandler struct {
	scheduleService services.SleepScheduleService
}

func NewSleepScheduleHandler(scheduleService services.SleepScheduleService) *SleepScheduleHandler {
	return &SleepScheduleHandler{
		scheduleService: scheduleService,
	}
}

// Get godoc
// @Summary Get sleep schedule
// @Description Get or create user's sleep schedule
// @Tags sleep-schedule
// @Produce json
// @Success 200 {object} dto.SleepScheduleResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/sleep-schedule [get]
func (h *SleepScheduleHandler) Get(c *gin.Context) {
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

	response, err := h.scheduleService.GetOrCreate(userID)
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

// Update godoc
// @Summary Update sleep schedule
// @Description Create or update user's sleep schedule
// @Tags sleep-schedule
// @Accept json
// @Produce json
// @Param request body dto.SleepScheduleRequest true "Sleep schedule request"
// @Success 200 {object} dto.SleepScheduleResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/sleep-schedule [put]
func (h *SleepScheduleHandler) Update(c *gin.Context) {
	var req dto.SleepScheduleRequest
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

	response, err := h.scheduleService.Update(userID, &req)
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
