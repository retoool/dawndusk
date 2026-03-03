package entities

import (
	"time"

	"github.com/google/uuid"
)

type Pet struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;unique" json:"user_id"`
	Name       string    `gorm:"type:varchar(50);not null" json:"name"`
	Type       string    `gorm:"type:varchar(50);not null" json:"type"` // cat, dog, bird, etc.
	Level      int       `gorm:"default:1" json:"level"`
	Experience int       `gorm:"default:0" json:"experience"`
	Health     int       `gorm:"default:100" json:"health"`
	Happiness  int       `gorm:"default:100" json:"happiness"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Pet) TableName() string {
	return "pets"
}

// CalculateLevel calculates the level based on experience points
func (p *Pet) CalculateLevel() int {
	// Simple formula: level = floor(sqrt(experience / 100)) + 1
	// Example: 0-99 exp = level 1, 100-399 exp = level 2, etc.
	if p.Experience < 100 {
		return 1
	}
	level := 1
	expRequired := 100
	totalExp := 0

	for totalExp+expRequired <= p.Experience {
		totalExp += expRequired
		level++
		expRequired += 50 // Each level requires 50 more exp
	}

	return level
}

// AddExperience adds experience points and updates level
func (p *Pet) AddExperience(exp int) {
	p.Experience += exp
	p.Level = p.CalculateLevel()
}

// GetExpForNextLevel returns experience needed for next level
func (p *Pet) GetExpForNextLevel() int {
	currentLevelExp := 100 + (p.Level-1)*50
	return currentLevelExp
}
