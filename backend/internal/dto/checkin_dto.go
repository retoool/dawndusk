package dto

import "time"

type CreateCheckInRequest struct {
	Type          string   `json:"type" binding:"required,oneof=wake sleep"`
	ScheduledTime string   `json:"scheduled_time" binding:"required"` // ISO 8601 format
	ActualTime    string   `json:"actual_time" binding:"required"`    // ISO 8601 format
	Mood          *string  `json:"mood,omitempty"`
	Note          *string  `json:"note,omitempty"`
	LocationLat   *float64 `json:"location_lat,omitempty"`
	LocationLng   *float64 `json:"location_lng,omitempty"`
}

type CheckInResponse struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Type           string    `json:"type"`
	ScheduledTime  time.Time `json:"scheduled_time"`
	ActualTime     time.Time `json:"actual_time"`
	TimeDifference *int      `json:"time_difference,omitempty"`
	Mood           *string   `json:"mood,omitempty"`
	Note           *string   `json:"note,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

type CheckInStatsResponse struct {
	TotalCheckIns     int     `json:"total_check_ins"`
	WakeCheckIns      int     `json:"wake_check_ins"`
	SleepCheckIns     int     `json:"sleep_check_ins"`
	CurrentStreak     int     `json:"current_streak"`
	LongestStreak     int     `json:"longest_streak"`
	AverageTimeDiff   float64 `json:"average_time_diff"`
	OnTimePercentage  float64 `json:"on_time_percentage"`
}
