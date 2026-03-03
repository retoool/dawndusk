package dto

import "time"

type AddFriendRequest struct {
	FriendID string `json:"friend_id" binding:"required"`
}

type FriendshipResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	FriendID   string    `json:"friend_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	AvatarURL  *string   `json:"avatar_url,omitempty"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
}

type FriendRequestResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	AvatarURL *string   `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
