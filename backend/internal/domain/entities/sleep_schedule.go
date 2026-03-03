package entities

import (
	"time"

	"github.com/google/uuid"
)

type SleepSchedule struct {
	ID                 uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID             uuid.UUID `gorm:"type:uuid;not null;unique" json:"user_id"`
	WakeTime           string    `gorm:"type:time;not null" json:"wake_time"`       // HH:MM format
	SleepTime          string    `gorm:"type:time;not null" json:"sleep_time"`      // HH:MM format
	AICallEnabled      bool      `gorm:"default:false" json:"ai_call_enabled"`
	AICallWakeOffset   int       `gorm:"default:0" json:"ai_call_wake_offset"`      // minutes before wake_time
	AICallSleepOffset  int       `gorm:"default:0" json:"ai_call_sleep_offset"`     // minutes before sleep_time
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (SleepSchedule) TableName() string {
	return "sleep_schedules"
}
