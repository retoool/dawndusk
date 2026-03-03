package services

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageService interface {
	SendMessage(senderID, receiverID uuid.UUID, req *dto.SendMessageRequest) (*dto.MessageResponse, error)
	GetConversation(userID, otherUserID uuid.UUID, limit, offset int) ([]*dto.MessageResponse, error)
	GetMessages(userID uuid.UUID, limit, offset int) ([]*dto.MessageResponse, error)
	MarkAsRead(userID, messageID uuid.UUID) error
	MarkConversationAsRead(userID, otherUserID uuid.UUID) error
	GetUnreadCount(userID uuid.UUID) (int64, error)
	DeleteMessage(userID, messageID uuid.UUID) error
}

type messageService struct {
	messageRepo repositories.MessageRepository
	userRepo    repositories.UserRepository
}

func NewMessageService(messageRepo repositories.MessageRepository, userRepo repositories.UserRepository) MessageService {
	return &messageService{
		messageRepo: messageRepo,
		userRepo:    userRepo,
	}
}

func (s *messageService) SendMessage(senderID, receiverID uuid.UUID, req *dto.SendMessageRequest) (*dto.MessageResponse, error) {
	// Verify receiver exists
	_, err := s.userRepo.FindByID(receiverID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError("USER_NOT_FOUND", "Receiver not found", 404)
		}
		return nil, errors.ErrInternalServer
	}

	// Create message
	message := &entities.Message{
		SenderID:    &senderID,
		ReceiverID:  receiverID,
		Content:     req.Content,
		MessageType: req.MessageType,
		IsRead:      false,
	}

	if err := s.messageRepo.Create(message); err != nil {
		return nil, errors.ErrInternalServer
	}

	// Load relations
	message, err = s.messageRepo.FindByID(message.ID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return s.toResponse(message), nil
}

func (s *messageService) GetConversation(userID, otherUserID uuid.UUID, limit, offset int) ([]*dto.MessageResponse, error) {
	messages, err := s.messageRepo.FindConversation(userID, otherUserID, limit, offset)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		responses[i] = s.toResponse(message)
	}

	return responses, nil
}

func (s *messageService) GetMessages(userID uuid.UUID, limit, offset int) ([]*dto.MessageResponse, error) {
	messages, err := s.messageRepo.FindUserMessages(userID, limit, offset)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		responses[i] = s.toResponse(message)
	}

	return responses, nil
}

func (s *messageService) MarkAsRead(userID, messageID uuid.UUID) error {
	// Verify message belongs to user
	message, err := s.messageRepo.FindByID(messageID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError("MESSAGE_NOT_FOUND", "Message not found", 404)
		}
		return errors.ErrInternalServer
	}

	if message.ReceiverID != userID {
		return errors.NewAppError("PERMISSION_DENIED", "You can only mark your own messages as read", 403)
	}

	return s.messageRepo.MarkAsRead(messageID)
}

func (s *messageService) MarkConversationAsRead(userID, otherUserID uuid.UUID) error {
	return s.messageRepo.MarkConversationAsRead(userID, otherUserID)
}

func (s *messageService) GetUnreadCount(userID uuid.UUID) (int64, error) {
	return s.messageRepo.GetUnreadCount(userID)
}

func (s *messageService) DeleteMessage(userID, messageID uuid.UUID) error {
	// Verify message belongs to user
	message, err := s.messageRepo.FindByID(messageID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError("MESSAGE_NOT_FOUND", "Message not found", 404)
		}
		return errors.ErrInternalServer
	}

	if message.SenderID == nil || *message.SenderID != userID {
		return errors.NewAppError("PERMISSION_DENIED", "You can only delete your own messages", 403)
	}

	return s.messageRepo.Delete(messageID)
}

func (s *messageService) toResponse(message *entities.Message) *dto.MessageResponse {
	response := &dto.MessageResponse{
		ID:          message.ID.String(),
		ReceiverID:  message.ReceiverID.String(),
		Content:     message.Content,
		MessageType: message.MessageType,
		IsRead:      message.IsRead,
		CreatedAt:   message.CreatedAt,
	}

	if message.SenderID != nil {
		senderID := message.SenderID.String()
		response.SenderID = &senderID
	}

	if message.Sender != nil {
		response.SenderUsername = &message.Sender.Username
	}

	if message.ReadAt != nil {
		response.ReadAt = message.ReadAt
	}

	return response
}
