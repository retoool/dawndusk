package services

import (
	"time"

	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SleepScheduleService interface {
	GetOrCreate(userID uuid.UUID) (*dto.SleepScheduleResponse, error)
	Update(userID uuid.UUID, req *dto.SleepScheduleRequest) (*dto.SleepScheduleResponse, error)
}

type sleepScheduleService struct {
	scheduleRepo repositories.SleepScheduleRepository
}

func NewSleepScheduleService(scheduleRepo repositories.SleepScheduleRepository) SleepScheduleService {
	return &sleepScheduleService{
		scheduleRepo: scheduleRepo,
	}
}

func (s *sleepScheduleService) GetOrCreate(userID uuid.UUID) (*dto.SleepScheduleResponse, error) {
	// Try to find existing schedule
	schedule, err := s.scheduleRepo.FindByUserID(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.ErrInternalServer
	}

	// Create default schedule if doesn't exist
	if err == gorm.ErrRecordNotFound {
		schedule = &entities.SleepSchedule{
			UserID:            userID,
			WakeTime:          "07:00",
			SleepTime:         "23:00",
			AICallEnabled:     false,
			AICallWakeOffset:  0,
			AICallSleepOffset: 0,
		}

		if err := s.scheduleRepo.Create(schedule); err != nil {
			return nil, errors.ErrInternalServer
		}
	}

	return s.toResponse(schedule), nil
}

func (s *sleepScheduleService) Update(userID uuid.UUID, req *dto.SleepScheduleRequest) (*dto.SleepScheduleResponse, error) {
	// Validate wake time format
	_, err := time.Parse("15:04", req.WakeTime)
	if err != nil {
		return nil, errors.NewAppError("INVALID_TIME", "Invalid wake time format. Use HH:MM", 400)
	}

	// Validate sleep time format
	_, err = time.Parse("15:04", req.SleepTime)
	if err != nil {
		return nil, errors.NewAppError("INVALID_TIME", "Invalid sleep time format. Use HH:MM", 400)
	}

	// Find existing schedule
	schedule, err := s.scheduleRepo.FindByUserID(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.ErrInternalServer
	}

	// Create or update
	if err == gorm.ErrRecordNotFound {
		schedule = &entities.SleepSchedule{
			UserID:            userID,
			WakeTime:          req.WakeTime,
			SleepTime:         req.SleepTime,
			AICallEnabled:     req.AICallEnabled,
			AICallWakeOffset:  req.AICallWakeOffset,
			AICallSleepOffset: req.AICallSleepOffset,
		}
		if err := s.scheduleRepo.Create(schedule); err != nil {
			return nil, errors.ErrInternalServer
		}
	} else {
		schedule.WakeTime = req.WakeTime
		schedule.SleepTime = req.SleepTime
		schedule.AICallEnabled = req.AICallEnabled
		schedule.AICallWakeOffset = req.AICallWakeOffset
		schedule.AICallSleepOffset = req.AICallSleepOffset

		if err := s.scheduleRepo.Update(schedule); err != nil {
			return nil, errors.ErrInternalServer
		}
	}

	return s.toResponse(schedule), nil
}

func (s *sleepScheduleService) toResponse(schedule *entities.SleepSchedule) *dto.SleepScheduleResponse {
	return &dto.SleepScheduleResponse{
		ID:                 schedule.ID.String(),
		UserID:             schedule.UserID.String(),
		WakeTime:           schedule.WakeTime,
		SleepTime:          schedule.SleepTime,
		AICallEnabled:      schedule.AICallEnabled,
		AICallWakeOffset:   schedule.AICallWakeOffset,
		AICallSleepOffset:  schedule.AICallSleepOffset,
		CreatedAt:          schedule.CreatedAt,
		UpdatedAt:          schedule.UpdatedAt,
	}
}
