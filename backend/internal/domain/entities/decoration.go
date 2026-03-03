package entities

import (
	"time"

	"github.com/google/uuid"
)

type PetDecoration struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	Description     string    `gorm:"type:text" json:"description"`
	Category        string    `gorm:"type:varchar(50);not null" json:"category"` // hat, accessory, background
	ImageURL        string    `gorm:"type:varchar(500);not null" json:"image_url"`
	UnlockLevel     int       `gorm:"default:1" json:"unlock_level"`
	UnlockCheckIns  int       `gorm:"default:0" json:"unlock_check_ins"`
	Rarity          string    `gorm:"type:varchar(20);default:'common'" json:"rarity"` // common, rare, epic, legendary
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (PetDecoration) TableName() string {
	return "pet_decorations"
}

type UserPetDecoration struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	DecorationID uuid.UUID `gorm:"type:uuid;not null" json:"decoration_id"`
	IsEquipped   bool      `gorm:"default:false" json:"is_equipped"`
	UnlockedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"unlocked_at"`

	// Relations
	User       User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Decoration PetDecoration `gorm:"foreignKey:DecorationID" json:"decoration,omitempty"`
}

func (UserPetDecoration) TableName() string {
	return "user_pet_decorations"
}
