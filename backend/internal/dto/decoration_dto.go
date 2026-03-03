package dto

import "time"

// Decoration DTOs
type DecorationResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	ImageURL       string    `json:"image_url"`
	UnlockLevel    int       `json:"unlock_level"`
	UnlockCheckIns int       `json:"unlock_check_ins"`
	Rarity         string    `json:"rarity"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserDecorationResponse struct {
	ID             string    `json:"id"`
	DecorationID   string    `json:"decoration_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	ImageURL       string    `json:"image_url"`
	Rarity         string    `json:"rarity"`
	IsEquipped     bool      `json:"is_equipped"`
	UnlockedAt     time.Time `json:"unlocked_at"`
}

type UnlockDecorationRequest struct {
	DecorationID string `json:"decoration_id" binding:"required"`
}
