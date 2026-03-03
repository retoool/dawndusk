package services

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetService interface {
	GetOrCreatePet(userID uuid.UUID) (*dto.PetResponse, error)
	UpdatePet(userID uuid.UUID, req *dto.UpdatePetRequest) (*dto.PetResponse, error)
	AddExperience(userID uuid.UUID, points int) (*dto.PetResponse, error)
	GetDecorations(userID uuid.UUID) ([]*dto.PetDecorationResponse, error)
	EquipDecoration(userID uuid.UUID, decorationID uuid.UUID) error
}

type petService struct {
	petRepo repositories.PetRepository
	db      *gorm.DB
}

func NewPetService(petRepo repositories.PetRepository, db *gorm.DB) PetService {
	return &petService{
		petRepo: petRepo,
		db:      db,
	}
}

func (s *petService) GetOrCreatePet(userID uuid.UUID) (*dto.PetResponse, error) {
	// Try to find existing pet
	pet, err := s.petRepo.FindByUserID(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.ErrInternalServer
	}

	// Create pet if doesn't exist
	if err == gorm.ErrRecordNotFound {
		pet = &entities.Pet{
			UserID:     userID,
			Name:       "小宠物",
			Type:       "cat",
			Level:      1,
			Experience: 0,
			Health:     100,
			Happiness:  100,
		}

		if err := s.petRepo.Create(pet); err != nil {
			return nil, errors.ErrInternalServer
		}
	}

	return s.petToResponse(pet), nil
}

func (s *petService) UpdatePet(userID uuid.UUID, req *dto.UpdatePetRequest) (*dto.PetResponse, error) {
	// Find pet
	pet, err := s.petRepo.FindByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError("PET_NOT_FOUND", "Pet not found", 404)
		}
		return nil, errors.ErrInternalServer
	}

	// Update name if provided
	if req.Name != nil {
		pet.Name = *req.Name
	}

	// Save changes
	if err := s.petRepo.Update(pet); err != nil {
		return nil, errors.ErrInternalServer
	}

	return s.petToResponse(pet), nil
}

func (s *petService) AddExperience(userID uuid.UUID, points int) (*dto.PetResponse, error) {
	// Get or create pet
	pet, err := s.petRepo.FindByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create pet if doesn't exist
			return s.GetOrCreatePet(userID)
		}
		return nil, errors.ErrInternalServer
	}

	// Add experience
	pet.Experience += points

	// Check for level up
	expNeeded := s.calculateExpForNextLevel(pet.Level)
	for pet.Experience >= expNeeded {
		pet.Level++
		pet.Experience -= expNeeded
		expNeeded = s.calculateExpForNextLevel(pet.Level)

		// Increase health and happiness on level up
		pet.Health = min(100, pet.Health+5)
		pet.Happiness = min(100, pet.Happiness+5)
	}

	// Save changes
	if err := s.petRepo.Update(pet); err != nil {
		return nil, errors.ErrInternalServer
	}

	return s.petToResponse(pet), nil
}

func (s *petService) GetDecorations(userID uuid.UUID) ([]*dto.PetDecorationResponse, error) {
	// Get pet to check level
	pet, err := s.petRepo.FindByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*dto.PetDecorationResponse{}, nil
		}
		return nil, errors.ErrInternalServer
	}

	// Get all decorations
	var decorations []entities.PetDecoration
	if err := s.db.Find(&decorations).Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	// Get user's owned decorations
	var userDecorations []entities.UserPetDecoration
	if err := s.db.Where("user_id = ?", userID).Find(&userDecorations).Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	// Create map for quick lookup
	ownedMap := make(map[uuid.UUID]bool)
	equippedMap := make(map[uuid.UUID]bool)
	for _, ud := range userDecorations {
		ownedMap[ud.DecorationID] = true
		if ud.IsEquipped {
			equippedMap[ud.DecorationID] = true
		}
	}

	// Convert to response DTOs
	responses := make([]*dto.PetDecorationResponse, 0)
	for _, decoration := range decorations {
		// Only show decorations that are unlocked or owned
		if decoration.UnlockLevel <= pet.Level || ownedMap[decoration.ID] {
			responses = append(responses, &dto.PetDecorationResponse{
				ID:          decoration.ID.String(),
				Name:        decoration.Name,
				Description: decoration.Description,
				Category:    decoration.Category,
				ImageURL:    decoration.ImageURL,
				UnlockLevel: decoration.UnlockLevel,
				Rarity:      decoration.Rarity,
				IsOwned:     ownedMap[decoration.ID],
				IsEquipped:  equippedMap[decoration.ID],
				CreatedAt:   decoration.CreatedAt,
			})
		}
	}

	return responses, nil
}

func (s *petService) EquipDecoration(userID uuid.UUID, decorationID uuid.UUID) error {
	// Verify user owns the decoration
	var userDecoration entities.UserPetDecoration
	err := s.db.Where("user_id = ? AND decoration_id = ?", userID, decorationID).First(&userDecoration).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError("DECORATION_NOT_OWNED", "You don't own this decoration", 403)
		}
		return errors.ErrInternalServer
	}

	// Get decoration to find category
	var decoration entities.PetDecoration
	if err := s.db.Where("id = ?", decorationID).First(&decoration).Error; err != nil {
		return errors.ErrInternalServer
	}

	// Unequip all decorations of the same category
	if err := s.db.Model(&entities.UserPetDecoration{}).
		Where("user_id = ? AND decoration_id IN (SELECT id FROM pet_decorations WHERE category = ?)", userID, decoration.Category).
		Update("is_equipped", false).Error; err != nil {
		return errors.ErrInternalServer
	}

	// Equip the new decoration
	userDecoration.IsEquipped = true
	if err := s.db.Save(&userDecoration).Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// Helper functions

func (s *petService) petToResponse(pet *entities.Pet) *dto.PetResponse {
	return &dto.PetResponse{
		ID:              pet.ID.String(),
		UserID:          pet.UserID.String(),
		Name:            pet.Name,
		Type:            pet.Type,
		Level:           pet.Level,
		Experience:      pet.Experience,
		Health:          pet.Health,
		Happiness:       pet.Happiness,
		ExpForNextLevel: s.calculateExpForNextLevel(pet.Level),
		CreatedAt:       pet.CreatedAt,
		UpdatedAt:       pet.UpdatedAt,
	}
}

func (s *petService) calculateExpForNextLevel(level int) int {
	// Formula: level * 100
	return level * 100
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
