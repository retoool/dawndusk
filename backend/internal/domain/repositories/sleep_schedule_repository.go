package repositories

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SleepScheduleRepository interface {
	Create(schedule *entities.SleepSchedule) error
	FindByUserID(userID uuid.UUID) (*entities.SleepSchedule, error)
	Update(schedule *entities.SleepSchedule) error
	Delete(userID uuid.UUID) error
}

type sleepScheduleRepository struct {
	db *gorm.DB
}

func NewSleepScheduleRepository(db *gorm.DB) SleepScheduleRepository {
	return &sleepScheduleRepository{db: db}
}

func (r *sleepScheduleRepository) Create(schedule *entities.SleepSchedule) error {
	return r.db.Create(schedule).Error
}

func (r *sleepScheduleRepository) FindByUserID(userID uuid.UUID) (*entities.SleepSchedule, error) {
	var schedule entities.SleepSchedule
	err := r.db.Where("user_id = ?", userID).First(&schedule).Error
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *sleepScheduleRepository) Update(schedule *entities.SleepSchedule) error {
	return r.db.Save(schedule).Error
}

func (r *sleepScheduleRepository) Delete(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&entities.SleepSchedule{}).Error
}
