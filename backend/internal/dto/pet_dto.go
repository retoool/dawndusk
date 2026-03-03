package dto

import "time"

type CreatePetRequest struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
	Type string `json:"type" binding:"required,oneof=cat dog bird rabbit hamster"`
}

type UpdatePetRequest struct {
	Name *string `json:"name,omitempty" binding:"omitempty,min=1,max=50"`
}

type PetResponse struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	Name              string    `json:"name"`
	Type              string    `json:"type"`
	Level             int       `json:"level"`
	Experience        int       `json:"experience"`
	Health            int       `json:"health"`
	Happiness         int       `json:"happiness"`
	ExpForNextLevel   int       `json:"exp_for_next_level"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type DecorationResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	ImageURL     string    `json:"image_url"`
	UnlockLevel  int       `json:"unlock_level"`
	Rarity       string    `json:"rarity"`
	IsOwned      bool      `json:"is_owned"`
	IsEquipped   bool      `json:"is_equipped"`
	CreatedAt    time.Time `json:"created_at"`
}
