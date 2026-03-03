package repositories

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DecorationRepository interface {
	// Decoration operations
	CreateDecoration(decoration *entities.PetDecoration) error
	FindDecorationByID(id uuid.UUID) (*entities.PetDecoration, error)
	ListDecorations(limit, offset int) ([]*entities.PetDecoration, error)
	UpdateDecoration(decoration *entities.PetDecoration) error
	DeleteDecoration(id uuid.UUID) error

	// User decoration operations
	UnlockDecoration(userID, decorationID uuid.UUID) error
	GetUserDecorations(userID uuid.UUID) ([]*entities.UserPetDecoration, error)
	EquipDecoration(userID, decorationID uuid.UUID) error
	UnequipAllDecorations(userID uuid.UUID) error
	IsDecorationUnlocked(userID, decorationID uuid.UUID) (bool, error)
}

type decorationRepository struct {
	db *gorm.DB
}

func NewDecorationRepository(db *gorm.DB) DecorationRepository {
	return &decorationRepository{db: db}
}

func (r *decorationRepository) CreateDecoration(decoration *entities.PetDecoration) error {
	return r.db.Create(decoration).Error
}

func (r *decorationRepository) FindDecorationByID(id uuid.UUID) (*entities.PetDecoration, error) {
	var decoration entities.PetDecoration
	err := r.db.First(&decoration, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &decoration, nil
}

func (r *decorationRepository) ListDecorations(limit, offset int) ([]*entities.PetDecoration, error) {
	var decorations []*entities.PetDecoration
	err := r.db.Limit(limit).Offset(offset).Order("unlock_level ASC, rarity ASC").Find(&decorations).Error
	return decorations, err
}

func (r *decorationRepository) UpdateDecoration(decoration *entities.PetDecoration) error {
	return r.db.Save(decoration).Error
}

func (r *decorationRepository) DeleteDecoration(id uuid.UUID) error {
	return r.db.Delete(&entities.PetDecoration{}, "id = ?", id).Error
}

func (r *decorationRepository) UnlockDecoration(userID, decorationID uuid.UUID) error {
	userDecoration := &entities.UserPetDecoration{
		UserID:       userID,
		DecorationID: decorationID,
		IsEquipped:   false,
	}
	return r.db.Create(userDecoration).Error
}

func (r *decorationRepository) GetUserDecorations(userID uuid.UUID) ([]*entities.UserPetDecoration, error) {
	var userDecorations []*entities.UserPetDecoration
	err := r.db.Preload("Decoration").Where("user_id = ?", userID).Find(&userDecorations).Error
	return userDecorations, err
}

func (r *decorationRepository) EquipDecoration(userID, decorationID uuid.UUID) error {
	return r.db.Model(&entities.UserPetDecoration{}).
		Where("user_id = ? AND decoration_id = ?", userID, decorationID).
		Update("is_equipped", true).Error
}

func (r *decorationRepository) UnequipAllDecorations(userID uuid.UUID) error {
	return r.db.Model(&entities.UserPetDecoration{}).
		Where("user_id = ?", userID).
		Update("is_equipped", false).Error
}

func (r *decorationRepository) IsDecorationUnlocked(userID, decorationID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&entities.UserPetDecoration{}).
		Where("user_id = ? AND decoration_id = ?", userID, decorationID).
		Count(&count).Error
	return count > 0, err
}
