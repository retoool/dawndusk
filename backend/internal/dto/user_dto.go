package dto

import "time"

type SleepScheduleRequest struct {
	WakeTime           string `json:"wake_time" binding:"required"`           // HH:MM format
	SleepTime          string `json:"sleep_time" binding:"required"`          // HH:MM format
	AICallEnabled      bool   `json:"ai_call_enabled"`
	AICallWakeOffset   int    `json:"ai_call_wake_offset"`   // Minutes before/after wake time
	AICallSleepOffset  int    `json:"ai_call_sleep_offset"`  // Minutes before/after sleep time
}

type SleepScheduleResponse struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	WakeTime           string    `json:"wake_time"`            // HH:MM format
	SleepTime          string    `json:"sleep_time"`           // HH:MM format
	AICallEnabled      bool      `json:"ai_call_enabled"`
	AICallWakeOffset   int       `json:"ai_call_wake_offset"`
	AICallSleepOffset  int       `json:"ai_call_sleep_offset"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type UpdateUserProfileRequest struct {
	Username    *string `json:"username,omitempty" binding:"omitempty,min=3,max=50"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
}

type UserProfileResponse struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber *string   `json:"phone_number,omitempty"`
	AvatarURL   *string   `json:"avatar_url,omitempty"`
	Timezone    string    `json:"timezone"`
	IsVerified  bool      `json:"is_verified"`
	CreatedAt   time.Time `json:"created_at"`
}
