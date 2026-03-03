package dto

import "time"

type SleepScheduleRequest struct {
	WakeTime           string `json:"wake_time" binding:"required"`
	SleepTime          string `json:"sleep_time" binding:"required"`
	AICallEnabled      bool   `json:"ai_call_enabled"`
	AICallWakeOffset   int    `json:"ai_call_wake_offset"`
	AICallSleepOffset  int    `json:"ai_call_sleep_offset"`
}

type SleepScheduleResponse struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	WakeTime           string    `json:"wake_time"`
	SleepTime          string    `json:"sleep_time"`
	AICallEnabled      bool      `json:"ai_call_enabled"`
	AICallWakeOffset   int       `json:"ai_call_wake_offset"`
	AICallSleepOffset  int       `json:"ai_call_sleep_offset"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
