package entities

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	SenderID    *uuid.UUID `gorm:"type:uuid" json:"sender_id,omitempty"`
	ReceiverID  uuid.UUID  `gorm:"type:uuid;not null" json:"receiver_id"`
	Content     string     `gorm:"type:text;not null" json:"content"`
	MessageType string     `gorm:"type:varchar(20);default:'text'" json:"message_type"` // text, image, system
	IsRead      bool       `gorm:"default:false" json:"is_read"`
	ReadAt      *time.Time `json:"read_at,omitempty"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relations
	Sender   *User `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	Receiver User  `gorm:"foreignKey:ReceiverID" json:"receiver,omitempty"`
}

func (Message) TableName() string {
	return "messages"
}
