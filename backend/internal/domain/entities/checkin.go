package entities

import (
	"time"

	"github.com/google/uuid"
)

type CheckIn struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Type           string    `gorm:"type:varchar(20);not null" json:"type"` // wake or sleep
	ScheduledTime  time.Time `gorm:"not null" json:"scheduled_time"`
	ActualTime     time.Time `gorm:"not null" json:"actual_time"`
	TimeDifference *int      `json:"time_difference,omitempty"` // minutes difference
	Mood           *string   `gorm:"type:varchar(20)" json:"mood,omitempty"`
	Note           *string   `gorm:"type:text" json:"note,omitempty"`
	LocationLat    *float64  `gorm:"type:decimal(10,8)" json:"location_lat,omitempty"`
	LocationLng    *float64  `gorm:"type:decimal(11,8)" json:"location_lng,omitempty"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (CheckIn) TableName() string {
	return "check_ins"
}
