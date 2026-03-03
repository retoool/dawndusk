package services

import (
	"time"

	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/google/uuid"
)

type CheckInService interface {
	CreateCheckIn(userID uuid.UUID, req *dto.CreateCheckInRequest) (*dto.CheckInResponse, error)
	GetCheckIns(userID uuid.UUID, limit, offset int) ([]*dto.CheckInResponse, error)
	GetTodayCheckIns(userID uuid.UUID) (*dto.TodayCheckInsResponse, error)
	GetCheckInStats(userID uuid.UUID, days int) (*dto.CheckInStatsResponse, error)
}

type checkInService struct {
	checkInRepo repositories.CheckInRepository
	petService  PetService
}

func NewCheckInService(checkInRepo repositories.CheckInRepository, petService PetService) CheckInService {
	return &checkInService{
		checkInRepo: checkInRepo,
		petService:  petService,
	}
}

func (s *checkInService) CreateCheckIn(userID uuid.UUID, req *dto.CreateCheckInRequest) (*dto.CheckInResponse, error) {
	// Parse times
	scheduledTime, err := time.Parse(time.RFC3339, req.ScheduledTime)
	if err != nil {
		return nil, errors.NewAppError("INVALID_TIME", "Invalid scheduled time format", 400)
	}

	actualTime, err := time.Parse(time.RFC3339, req.ActualTime)
	if err != nil {
		return nil, errors.NewAppError("INVALID_TIME", "Invalid actual time format", 400)
	}

	// Calculate time difference in minutes
	timeDiff := int(actualTime.Sub(scheduledTime).Minutes())

	// Create check-in entity
	checkIn := &entities.CheckIn{
		UserID:         userID,
		Type:           req.Type,
		ScheduledTime:  scheduledTime,
		ActualTime:     actualTime,
		TimeDifference: &timeDiff,
		Mood:           req.Mood,
		Note:           req.Note,
		LocationLat:    req.LocationLat,
		LocationLng:    req.LocationLng,
	}

	// Save to database
	if err := s.checkInRepo.Create(checkIn); err != nil {
		return nil, errors.ErrInternalServer
	}

	// Calculate experience points based on punctuality
	expPoints := s.calculateExperiencePoints(timeDiff)

	// Update pet experience
	if _, err := s.petService.AddExperience(userID, expPoints); err != nil {
		// Log error but don't fail the check-in
		// In production, use proper logging
	}

	// Return response
	return &dto.CheckInResponse{
		ID:             checkIn.ID.String(),
		UserID:         checkIn.UserID.String(),
		Type:           checkIn.Type,
		ScheduledTime:  checkIn.ScheduledTime,
		ActualTime:     checkIn.ActualTime,
		TimeDifference: checkIn.TimeDifference,
		Mood:           checkIn.Mood,
		Note:           checkIn.Note,
		CreatedAt:      checkIn.CreatedAt,
	}, nil
}

func (s *checkInService) GetCheckIns(userID uuid.UUID, limit, offset int) ([]*dto.CheckInResponse, error) {
	// Set default limit if not provided
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	checkIns, err := s.checkInRepo.FindByUserID(userID, limit, offset)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	// Convert to response DTOs
	responses := make([]*dto.CheckInResponse, len(checkIns))
	for i, checkIn := range checkIns {
		responses[i] = &dto.CheckInResponse{
			ID:             checkIn.ID.String(),
			UserID:         checkIn.UserID.String(),
			Type:           checkIn.Type,
			ScheduledTime:  checkIn.ScheduledTime,
			ActualTime:     checkIn.ActualTime,
			TimeDifference: checkIn.TimeDifference,
			Mood:           checkIn.Mood,
			Note:           checkIn.Note,
			CreatedAt:      checkIn.CreatedAt,
		}
	}

	return responses, nil
}

func (s *checkInService) GetTodayCheckIns(userID uuid.UUID) (*dto.TodayCheckInsResponse, error) {
	checkIns, err := s.checkInRepo.FindTodayByUserID(userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := &dto.TodayCheckInsResponse{
		HasWake:  false,
		HasSleep: false,
	}

	// Separate wake and sleep check-ins
	for _, checkIn := range checkIns {
		checkInResp := &dto.CheckInResponse{
			ID:             checkIn.ID.String(),
			UserID:         checkIn.UserID.String(),
			Type:           checkIn.Type,
			ScheduledTime:  checkIn.ScheduledTime,
			ActualTime:     checkIn.ActualTime,
			TimeDifference: checkIn.TimeDifference,
			Mood:           checkIn.Mood,
			Note:           checkIn.Note,
			CreatedAt:      checkIn.CreatedAt,
		}

		if checkIn.Type == "wake" {
			response.WakeCheckIn = checkInResp
			response.HasWake = true
		} else if checkIn.Type == "sleep" {
			response.SleepCheckIn = checkInResp
			response.HasSleep = true
		}
	}

	return response, nil
}

func (s *checkInService) GetCheckInStats(userID uuid.UUID, days int) (*dto.CheckInStatsResponse, error) {
	if days <= 0 {
		days = 30
	}

	// Get all check-ins for the user (we'll filter by date in memory for simplicity)
	// In production, you'd want to add a date filter to the repository method
	checkIns, err := s.checkInRepo.FindByUserID(userID, 1000, 0)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	// Calculate statistics
	stats := &dto.CheckInStatsResponse{
		TotalCheckIns:    len(checkIns),
		WakeCheckIns:     0,
		SleepCheckIns:    0,
		CurrentStreak:    0,
		LongestStreak:    0,
		AverageTimeDiff:  0,
		OnTimePercentage: 0,
	}

	if len(checkIns) == 0 {
		return stats, nil
	}

	// Count wake and sleep check-ins
	var totalTimeDiff int
	var onTimeCount int
	for _, checkIn := range checkIns {
		if checkIn.Type == "wake" {
			stats.WakeCheckIns++
		} else if checkIn.Type == "sleep" {
			stats.SleepCheckIns++
		}

		if checkIn.TimeDifference != nil {
			totalTimeDiff += *checkIn.TimeDifference
			// Consider on-time if within 15 minutes
			if *checkIn.TimeDifference >= -15 && *checkIn.TimeDifference <= 15 {
				onTimeCount++
			}
		}
	}

	// Calculate average time difference
	if stats.TotalCheckIns > 0 {
		stats.AverageTimeDiff = float64(totalTimeDiff) / float64(stats.TotalCheckIns)
		stats.OnTimePercentage = float64(onTimeCount) / float64(stats.TotalCheckIns) * 100
	}

	// Calculate streaks
	stats.CurrentStreak, stats.LongestStreak = s.calculateStreaks(checkIns)

	return stats, nil
}

// Helper function to calculate experience points based on punctuality
func (s *checkInService) calculateExperiencePoints(timeDiffMinutes int) int {
	// Base points
	basePoints := 5

	// Bonus for being on time (within 15 minutes)
	if timeDiffMinutes >= -15 && timeDiffMinutes <= 15 {
		return basePoints + 5 // 10 points total
	}

	// Penalty for being very late (more than 1 hour)
	if timeDiffMinutes > 60 || timeDiffMinutes < -60 {
		return basePoints - 2 // 3 points
	}

	return basePoints
}

// Helper function to calculate current and longest streaks
func (s *checkInService) calculateStreaks(checkIns []entities.CheckIn) (int, int) {
	if len(checkIns) == 0 {
		return 0, 0
	}

	// Group check-ins by date
	dateMap := make(map[string]bool)
	for _, checkIn := range checkIns {
		date := checkIn.CreatedAt.Format("2006-01-02")
		dateMap[date] = true
	}

	// Convert to sorted slice of dates
	var dates []time.Time
	for dateStr := range dateMap {
		date, _ := time.Parse("2006-01-02", dateStr)
		dates = append(dates, date)
	}

	// Sort dates in descending order
	for i := 0; i < len(dates)-1; i++ {
		for j := i + 1; j < len(dates); j++ {
			if dates[i].Before(dates[j]) {
				dates[i], dates[j] = dates[j], dates[i]
			}
		}
	}

	// Calculate current streak
	currentStreak := 0
	today := time.Now().Truncate(24 * time.Hour)
	for i, date := range dates {
		expectedDate := today.AddDate(0, 0, -i)
		if date.Equal(expectedDate) {
			currentStreak++
		} else {
			break
		}
	}

	// Calculate longest streak
	longestStreak := 0
	tempStreak := 1
	for i := 1; i < len(dates); i++ {
		diff := dates[i-1].Sub(dates[i]).Hours() / 24
		if diff == 1 {
			tempStreak++
			if tempStreak > longestStreak {
				longestStreak = tempStreak
			}
		} else {
			tempStreak = 1
		}
	}
	if tempStreak > longestStreak {
		longestStreak = tempStreak
	}

	return currentStreak, longestStreak
}
