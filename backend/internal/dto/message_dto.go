package dto

import "time"

// Message DTOs
type SendMessageRequest struct {
	ReceiverID  string `json:"receiver_id" binding:"required"`
	Content     string `json:"content" binding:"required,min=1,max=1000"`
	MessageType string `json:"message_type" binding:"omitempty,oneof=text image system"`
}

type MessageResponse struct {
	ID             string     `json:"id"`
	SenderID       *string    `json:"sender_id,omitempty"`
	SenderUsername *string    `json:"sender_username,omitempty"`
	ReceiverID     string     `json:"receiver_id"`
	Content        string     `json:"content"`
	MessageType    string     `json:"message_type"`
	IsRead         bool       `json:"is_read"`
	ReadAt         *time.Time `json:"read_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

type UnreadCountResponse struct {
	Count int64 `json:"count"`
}
