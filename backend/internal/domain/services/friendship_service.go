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

type FriendshipService interface {
	SendFriendRequest(userID, friendID uuid.UUID) (*dto.FriendshipResponse, error)
	AcceptFriendRequest(userID, requestID uuid.UUID) error
	RejectFriendRequest(userID, requestID uuid.UUID) error
	RemoveFriend(userID, friendshipID uuid.UUID) error
	GetFriends(userID uuid.UUID) ([]*dto.FriendshipResponse, error)
	GetPendingRequests(userID uuid.UUID) ([]*dto.FriendRequestResponse, error)
	GetSentRequests(userID uuid.UUID) ([]*dto.FriendRequestResponse, error)
}

type friendshipService struct {
	friendshipRepo repositories.FriendshipRepository
	userRepo       repositories.UserRepository
}

func NewFriendshipService(friendshipRepo repositories.FriendshipRepository, userRepo repositories.UserRepository) FriendshipService {
	return &friendshipService{
		friendshipRepo: friendshipRepo,
		userRepo:       userRepo,
	}
}

func (s *friendshipService) SendFriendRequest(userID, friendID uuid.UUID) (*dto.FriendshipResponse, error) {
	// Check if trying to add self
	if userID == friendID {
		return nil, errors.NewAppError("CANNOT_ADD_SELF", "Cannot add yourself as friend", 400)
	}

	// Check if friend exists
	friend, err := s.userRepo.FindByID(friendID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError("USER_NOT_FOUND", "User not found", 404)
		}
		return nil, errors.ErrInternalServer
	}

	// Check if friendship already exists
	existing, err := s.friendshipRepo.FindByUserAndFriend(userID, friendID)
	if err == nil && existing != nil {
		if existing.Status == "accepted" {
			return nil, errors.NewAppError("ALREADY_FRIENDS", "Already friends", 400)
		}
		if existing.Status == "pending" {
			return nil, errors.NewAppError("REQUEST_PENDING", "Friend request already sent", 400)
		}
	}

	// Create friendship
	friendship := &entities.Friendship{
		UserID:   userID,
		FriendID: friendID,
		Status:   "pending",
	}

	if err := s.friendshipRepo.Create(friendship); err != nil {
		return nil, errors.ErrInternalServer
	}

	return &dto.FriendshipResponse{
		ID:        friendship.ID.String(),
		UserID:    friendship.UserID.String(),
		FriendID:  friendship.FriendID.String(),
		Username:  friend.Username,
		Email:     friend.Email,
		AvatarURL: friend.AvatarURL,
		Status:    friendship.Status,
		CreatedAt: friendship.CreatedAt,
	}, nil
}

func (s *friendshipService) AcceptFriendRequest(userID, requestID uuid.UUID) error {
	// Find the request
	friendship, err := s.friendshipRepo.FindByID(requestID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError("REQUEST_NOT_FOUND", "Friend request not found", 404)
		}
		return errors.ErrInternalServer
	}

	// Verify the request is for this user
	if friendship.FriendID != userID {
		return errors.NewAppError("PERMISSION_DENIED", "You can only accept requests sent to you", 403)
	}

	// Check if already accepted
	if friendship.Status == "accepted" {
		return errors.NewAppError("ALREADY_ACCEPTED", "Request already accepted", 400)
	}

	// Update status
	now := time.Now()
	friendship.Status = "accepted"
	friendship.AcceptedAt = &now

	if err := s.friendshipRepo.UpdateStatus(requestID, "accepted"); err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

func (s *friendshipService) RejectFriendRequest(userID, requestID uuid.UUID) error {
	// Find the request
	friendship, err := s.friendshipRepo.FindByID(requestID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError("REQUEST_NOT_FOUND", "Friend request not found", 404)
		}
		return errors.ErrInternalServer
	}

	// Verify the request is for this user
	if friendship.FriendID != userID {
		return errors.NewAppError("PERMISSION_DENIED", "You can only reject requests sent to you", 403)
	}

	// Delete the request
	return s.friendshipRepo.Delete(requestID)
}

func (s *friendshipService) RemoveFriend(userID, friendshipID uuid.UUID) error {
	// Find the friendship
	friendship, err := s.friendshipRepo.FindByID(friendshipID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError("FRIENDSHIP_NOT_FOUND", "Friendship not found", 404)
		}
		return errors.ErrInternalServer
	}

	// Verify user is part of this friendship
	if friendship.UserID != userID && friendship.FriendID != userID {
		return errors.NewAppError("PERMISSION_DENIED", "You can only remove your own friendships", 403)
	}

	// Delete the friendship
	return s.friendshipRepo.Delete(friendshipID)
}

func (s *friendshipService) GetFriends(userID uuid.UUID) ([]*dto.FriendshipResponse, error) {
	friendships, err := s.friendshipRepo.GetFriends(userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.FriendshipResponse, len(friendships))
	for i, friendship := range friendships {
		// Determine which user is the friend
		var friend entities.User
		if friendship.UserID == userID {
			friend = friendship.Friend
		} else {
			friend = friendship.User
		}

		responses[i] = &dto.FriendshipResponse{
			ID:         friendship.ID.String(),
			UserID:     friendship.UserID.String(),
			FriendID:   friendship.FriendID.String(),
			Username:   friend.Username,
			Email:      friend.Email,
			AvatarURL:  friend.AvatarURL,
			Status:     friendship.Status,
			CreatedAt:  friendship.CreatedAt,
			AcceptedAt: friendship.AcceptedAt,
		}
	}

	return responses, nil
}

func (s *friendshipService) GetPendingRequests(userID uuid.UUID) ([]*dto.FriendRequestResponse, error) {
	friendships, err := s.friendshipRepo.GetPendingRequests(userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.FriendRequestResponse, len(friendships))
	for i, friendship := range friendships {
		responses[i] = &dto.FriendRequestResponse{
			ID:        friendship.ID.String(),
			UserID:    friendship.User.ID.String(),
			Username:  friendship.User.Username,
			Email:     friendship.User.Email,
			AvatarURL: friendship.User.AvatarURL,
			CreatedAt: friendship.CreatedAt,
		}
	}

	return responses, nil
}

func (s *friendshipService) GetSentRequests(userID uuid.UUID) ([]*dto.FriendRequestResponse, error) {
	friendships, err := s.friendshipRepo.GetSentRequests(userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.FriendRequestResponse, len(friendships))
	for i, friendship := range friendships {
		responses[i] = &dto.FriendRequestResponse{
			ID:        friendship.ID.String(),
			UserID:    friendship.Friend.ID.String(),
			Username:  friendship.Friend.Username,
			Email:     friendship.Friend.Email,
			AvatarURL: friendship.Friend.AvatarURL,
			CreatedAt: friendship.CreatedAt,
		}
	}

	return responses, nil
}
