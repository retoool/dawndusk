package services

import (
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DecorationService interface {
	ListDecorations(limit, offset int) ([]*dto.DecorationResponse, error)
	GetUserDecorations(userID uuid.UUID) ([]*dto.UserDecorationResponse, error)
	UnlockDecoration(userID, decorationID uuid.UUID) error
	EquipDecoration(userID, decorationID uuid.UUID) error
}

type decorationService struct {
	decorationRepo repositories.DecorationRepository
	petRepo        repositories.PetRepository
}

func NewDecorationService(decorationRepo repositories.DecorationRepository, petRepo repositories.PetRepository) DecorationService {
	return &decorationService{
		decorationRepo: decorationRepo,
		petRepo:        petRepo,
	}
}

func (s *decorationService) ListDecorations(limit, offset int) ([]*dto.DecorationResponse, error) {
	decorations, err := s.decorationRepo.ListDecorations(limit, offset)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.DecorationResponse, len(decorations))
	for i, decoration := range decorations {
		responses[i] = &dto.DecorationResponse{
			ID:             decoration.ID.String(),
			Name:           decoration.Name,
			Description:    decoration.Description,
			Category:       decoration.Category,
			ImageURL:       decoration.ImageURL,
			UnlockLevel:    decoration.UnlockLevel,
			UnlockCheckIns: decoration.UnlockCheckIns,
			Rarity:         decoration.Rarity,
			CreatedAt:      decoration.CreatedAt,
		}
	}

	return responses, nil
}

func (s *decorationService) GetUserDecorations(userID uuid.UUID) ([]*dto.UserDecorationResponse, error) {
	userDecorations, err := s.decorationRepo.GetUserDecorations(userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.UserDecorationResponse, len(userDecorations))
	for i, userDecoration := range userDecorations {
		responses[i] = &dto.UserDecorationResponse{
			ID:             userDecoration.ID.String(),
			DecorationID:   userDecoration.DecorationID.String(),
			Name:           userDecoration.Decoration.Name,
			Description:    userDecoration.Decoration.Description,
			Category:       userDecoration.Decoration.Category,
			ImageURL:       userDecoration.Decoration.ImageURL,
			Rarity:         userDecoration.Decoration.Rarity,
			IsEquipped:     userDecoration.IsEquipped,
			UnlockedAt:     userDecoration.UnlockedAt,
		}
	}

	return responses, nil
}

func (s *decorationService) UnlockDecoration(userID, decorationID uuid.UUID) error {
	// Check if decoration exists
	decoration, err := s.decorationRepo.FindDecorationByID(decorationID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError("DECORATION_NOT_FOUND", "Decoration not found", 404)
		}
		return errors.ErrInternalServer
	}

	// Check if already unlocked
	isUnlocked, err := s.decorationRepo.IsDecorationUnlocked(userID, decorationID)
	if err != nil {
		return errors.ErrInternalServer
	}

	if isUnlocked {
		return errors.NewAppError("ALREADY_UNLOCKED", "Decoration already unlocked", 400)
	}

	// Check if user meets requirements
	pet, err := s.petRepo.FindByUserID(userID)
	if err != nil {
		return errors.ErrInternalServer
	}

	if pet.Level < decoration.UnlockLevel {
		return errors.NewAppError("LEVEL_TOO_LOW", "Pet level too low to unlock this decoration", 400)
	}

	// Unlock decoration
	return s.decorationRepo.UnlockDecoration(userID, decorationID)
}

func (s *decorationService) EquipDecoration(userID, decorationID uuid.UUID) error {
	// Check if decoration is unlocked
	isUnlocked, err := s.decorationRepo.IsDecorationUnlocked(userID, decorationID)
	if err != nil {
		return errors.ErrInternalServer
	}

	if !isUnlocked {
		return errors.NewAppError("NOT_UNLOCKED", "Decoration not unlocked", 400)
	}

	// Unequip all decorations first
	if err := s.decorationRepo.UnequipAllDecorations(userID); err != nil {
		return errors.ErrInternalServer
	}

	// Equip the decoration
	return s.decorationRepo.EquipDecoration(userID, decorationID)
}
