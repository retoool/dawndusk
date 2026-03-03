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

type MessageHandler struct {
	messageService services.MessageService
}

func NewMessageHandler(messageService services.MessageService) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
	}
}

// SendMessage godoc
// @Summary Send a message
// @Description Send a message to another user
// @Tags messages
// @Accept json
// @Produce json
// @Param request body dto.SendMessageRequest true "Send message request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/messages [post]
func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req dto.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDStr, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	senderID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	receiverID, err := uuid.Parse(req.ReceiverID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver ID"})
		return
	}

	response, err := h.messageService.SendMessage(senderID, receiverID, &req)
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

// GetConversation godoc
// @Summary Get conversation with a user
// @Description Get all messages in a conversation with another user
// @Tags messages
// @Produce json
// @Param userId path string true "Other user ID"
// @Param limit query int false "Limit" default(50)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.MessageResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/messages/conversation/{userId} [get]
func (h *MessageHandler) GetConversation(c *gin.Context) {
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

	otherUserID, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	responses, err := h.messageService.GetConversation(userID, otherUserID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// GetMessages godoc
// @Summary Get user messages
// @Description Get all messages for the current user
// @Tags messages
// @Produce json
// @Param limit query int false "Limit" default(50)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.MessageResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/messages [get]
func (h *MessageHandler) GetMessages(c *gin.Context) {
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

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	responses, err := h.messageService.GetMessages(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// MarkAsRead godoc
// @Summary Mark message as read
// @Description Mark a specific message as read
// @Tags messages
// @Param id path string true "Message ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/messages/{id}/read [post]
func (h *MessageHandler) MarkAsRead(c *gin.Context) {
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

	messageID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	if err := h.messageService.MarkAsRead(userID, messageID); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message marked as read"})
}

// MarkConversationAsRead godoc
// @Summary Mark conversation as read
// @Description Mark all messages in a conversation as read
// @Tags messages
// @Param userId path string true "Other user ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/messages/conversation/{userId}/read [post]
func (h *MessageHandler) MarkConversationAsRead(c *gin.Context) {
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

	otherUserID, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.messageService.MarkConversationAsRead(userID, otherUserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversation marked as read"})
}

// GetUnreadCount godoc
// @Summary Get unread message count
// @Description Get the count of unread messages for the current user
// @Tags messages
// @Produce json
// @Success 200 {object} dto.UnreadCountResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/messages/unread-count [get]
func (h *MessageHandler) GetUnreadCount(c *gin.Context) {
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

	count, err := h.messageService.GetUnreadCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.UnreadCountResponse{Count: count})
}

// DeleteMessage godoc
// @Summary Delete a message
// @Description Delete a message (sender only)
// @Tags messages
// @Param id path string true "Message ID"
// @Success 204
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/messages/{id} [delete]
func (h *MessageHandler) DeleteMessage(c *gin.Context) {
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

	messageID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	if err := h.messageService.DeleteMessage(userID, messageID); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message, "code": appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusNoContent)
}
