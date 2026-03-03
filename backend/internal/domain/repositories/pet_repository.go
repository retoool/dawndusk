package repositories

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetRepository interface {
	Create(pet *entities.Pet) error
	FindByID(id uuid.UUID) (*entities.Pet, error)
	FindByUserID(userID uuid.UUID) (*entities.Pet, error)
	Update(pet *entities.Pet) error
	Delete(id uuid.UUID) error
}

type petRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) PetRepository {
	return &petRepository{db: db}
}

func (r *petRepository) Create(pet *entities.Pet) error {
	return r.db.Create(pet).Error
}

func (r *petRepository) FindByID(id uuid.UUID) (*entities.Pet, error) {
	var pet entities.Pet
	err := r.db.Where("id = ?", id).First(&pet).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (r *petRepository) FindByUserID(userID uuid.UUID) (*entities.Pet, error) {
	var pet entities.Pet
	err := r.db.Where("user_id = ?", userID).First(&pet).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (r *petRepository) Update(pet *entities.Pet) error {
	return r.db.Save(pet).Error
}

func (r *petRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Pet{}, id).Error
}
