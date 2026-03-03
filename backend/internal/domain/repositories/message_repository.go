package repositories

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(message *entities.Message) error
	FindByID(id uuid.UUID) (*entities.Message, error)
	FindConversation(userID1, userID2 uuid.UUID, limit, offset int) ([]*entities.Message, error)
	FindUserMessages(userID uuid.UUID, limit, offset int) ([]*entities.Message, error)
	MarkAsRead(messageID uuid.UUID) error
	MarkConversationAsRead(receiverID, senderID uuid.UUID) error
	GetUnreadCount(userID uuid.UUID) (int64, error)
	Delete(id uuid.UUID) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *entities.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) FindByID(id uuid.UUID) (*entities.Message, error) {
	var message entities.Message
	err := r.db.Preload("Sender").Preload("Receiver").First(&message, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *messageRepository) FindConversation(userID1, userID2 uuid.UUID, limit, offset int) ([]*entities.Message, error) {
	var messages []*entities.Message
	err := r.db.Preload("Sender").Preload("Receiver").
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			userID1, userID2, userID2, userID1).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, err
}

func (r *messageRepository) FindUserMessages(userID uuid.UUID, limit, offset int) ([]*entities.Message, error) {
	var messages []*entities.Message
	err := r.db.Preload("Sender").Preload("Receiver").
		Where("receiver_id = ? OR sender_id = ?", userID, userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, err
}

func (r *messageRepository) MarkAsRead(messageID uuid.UUID) error {
	now := gorm.Expr("CURRENT_TIMESTAMP")
	return r.db.Model(&entities.Message{}).
		Where("id = ?", messageID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": now,
		}).Error
}

func (r *messageRepository) MarkConversationAsRead(receiverID, senderID uuid.UUID) error {
	now := gorm.Expr("CURRENT_TIMESTAMP")
	return r.db.Model(&entities.Message{}).
		Where("receiver_id = ? AND sender_id = ? AND is_read = false", receiverID, senderID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": now,
		}).Error
}

func (r *messageRepository) GetUnreadCount(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Message{}).
		Where("receiver_id = ? AND is_read = false", userID).
		Count(&count).Error
	return count, err
}

func (r *messageRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Message{}, "id = ?", id).Error
}
