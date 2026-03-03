package entities

import (
	"time"

	"github.com/google/uuid"
)

type Friendship struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index:idx_friendship_user" json:"user_id"`
	FriendID   uuid.UUID `gorm:"type:uuid;not null;index:idx_friendship_friend" json:"friend_id"`
	Status     string    `gorm:"type:varchar(20);not null;default:'pending'" json:"status"` // pending, accepted, blocked
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`

	// Relations
	User   User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Friend User `gorm:"foreignKey:FriendID" json:"friend,omitempty"`
}

func (Friendship) TableName() string {
	return "friendships"
}
