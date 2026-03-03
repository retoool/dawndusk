package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username     string     `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email        string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string     `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber  *string    `gorm:"type:varchar(20)" json:"phone_number,omitempty"`
	AvatarURL    *string    `gorm:"type:varchar(500)" json:"avatar_url,omitempty"`
	Timezone     string     `gorm:"type:varchar(50);default:'UTC'" json:"timezone"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
	IsActive     bool       `gorm:"default:true" json:"is_active"`
	IsVerified   bool       `gorm:"default:false" json:"is_verified"`
}

func (User) TableName() string {
	return "users"
}
