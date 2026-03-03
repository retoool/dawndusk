package entities

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Description *string    `gorm:"type:text" json:"description,omitempty"`
	AvatarURL   *string    `gorm:"type:varchar(500)" json:"avatar_url,omitempty"`
	CreatorID   *uuid.UUID `gorm:"type:uuid" json:"creator_id,omitempty"`
	MaxMembers  int        `gorm:"default:50" json:"max_members"`
	IsPrivate   bool       `gorm:"default:false" json:"is_private"`
	InviteCode  string     `gorm:"type:varchar(20);unique" json:"invite_code"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Creator *User `gorm:"foreignKey:CreatorID" json:"creator,omitempty"`
}

func (Group) TableName() string {
	return "groups"
}

type GroupMember struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	GroupID   uuid.UUID `gorm:"type:uuid;not null" json:"group_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Role      string    `gorm:"type:varchar(20);default:'member'" json:"role"` // admin, moderator, member
	JoinedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`

	// Relations
	Group Group `gorm:"foreignKey:GroupID" json:"group,omitempty"`
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (GroupMember) TableName() string {
	return "group_members"
}
