package dto

import "time"

// Group DTOs
type CreateGroupRequest struct {
	Name        string  `json:"name" binding:"required,min=2,max=100"`
	Description *string `json:"description,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	MaxMembers  int     `json:"max_members" binding:"required,min=2,max=500"`
	IsPrivate   bool    `json:"is_private"`
}

type UpdateGroupRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=2,max=100"`
	Description *string `json:"description,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	MaxMembers  *int    `json:"max_members,omitempty" binding:"omitempty,min=2,max=500"`
	IsPrivate   *bool   `json:"is_private,omitempty"`
}

type GroupResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	AvatarURL   *string   `json:"avatar_url,omitempty"`
	CreatorID   *string   `json:"creator_id,omitempty"`
	MaxMembers  int       `json:"max_members"`
	IsPrivate   bool      `json:"is_private"`
	InviteCode  string    `json:"invite_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JoinGroupRequest struct {
	InviteCode string `json:"invite_code" binding:"required"`
}

type GroupMemberResponse struct {
	ID       string    `json:"id"`
	GroupID  string    `json:"group_id"`
	UserID   string    `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}

type UpdateMemberRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=admin moderator member"`
}
