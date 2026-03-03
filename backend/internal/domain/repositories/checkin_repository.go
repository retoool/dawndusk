package repositories

import (
	"time"

	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckInRepository interface {
	Create(checkIn *entities.CheckIn) error
	FindByID(id uuid.UUID) (*entities.CheckIn, error)
	FindByUserID(userID uuid.UUID, limit, offset int) ([]entities.CheckIn, error)
	FindTodayByUserID(userID uuid.UUID) ([]entities.CheckIn, error)
	CountByUserID(userID uuid.UUID) (int64, error)
	Delete(id uuid.UUID) error
}

type checkInRepository struct {
	db *gorm.DB
}

func NewCheckInRepository(db *gorm.DB) CheckInRepository {
	return &checkInRepository{db: db}
}

func (r *checkInRepository) Create(checkIn *entities.CheckIn) error {
	return r.db.Create(checkIn).Error
}

func (r *checkInRepository) FindByID(id uuid.UUID) (*entities.CheckIn, error) {
	var checkIn entities.CheckIn
	err := r.db.Where("id = ?", id).First(&checkIn).Error
	if err != nil {
		return nil, err
	}
	return &checkIn, nil
}

func (r *checkInRepository) FindByUserID(userID uuid.UUID, limit, offset int) ([]entities.CheckIn, error) {
	var checkIns []entities.CheckIn
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&checkIns).Error
	return checkIns, err
}

func (r *checkInRepository) FindTodayByUserID(userID uuid.UUID) ([]entities.CheckIn, error) {
	var checkIns []entities.CheckIn
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	err := r.db.Where("user_id = ? AND created_at >= ? AND created_at < ?", userID, today, tomorrow).
		Order("created_at DESC").
		Find(&checkIns).Error
	return checkIns, err
}

func (r *checkInRepository) CountByUserID(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&entities.CheckIn{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *checkInRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.CheckIn{}, id).Error
}
