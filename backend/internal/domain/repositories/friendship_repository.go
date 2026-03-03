package repositories

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FriendshipRepository interface {
	Create(friendship *entities.Friendship) error
	FindByID(id uuid.UUID) (*entities.Friendship, error)
	FindByUserAndFriend(userID, friendID uuid.UUID) (*entities.Friendship, error)
	GetFriends(userID uuid.UUID) ([]*entities.Friendship, error)
	GetPendingRequests(userID uuid.UUID) ([]*entities.Friendship, error)
	GetSentRequests(userID uuid.UUID) ([]*entities.Friendship, error)
	UpdateStatus(id uuid.UUID, status string) error
	Delete(id uuid.UUID) error
	AreFriends(userID, friendID uuid.UUID) (bool, error)
}

type friendshipRepository struct {
	db *gorm.DB
}

func NewFriendshipRepository(db *gorm.DB) FriendshipRepository {
	return &friendshipRepository{db: db}
}

func (r *friendshipRepository) Create(friendship *entities.Friendship) error {
	return r.db.Create(friendship).Error
}

func (r *friendshipRepository) FindByID(id uuid.UUID) (*entities.Friendship, error) {
	var friendship entities.Friendship
	err := r.db.Preload("User").Preload("Friend").First(&friendship, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &friendship, nil
}

func (r *friendshipRepository) FindByUserAndFriend(userID, friendID uuid.UUID) (*entities.Friendship, error) {
	var friendship entities.Friendship
	err := r.db.Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
		userID, friendID, friendID, userID).First(&friendship).Error
	if err != nil {
		return nil, err
	}
	return &friendship, nil
}

func (r *friendshipRepository) GetFriends(userID uuid.UUID) ([]*entities.Friendship, error) {
	var friendships []*entities.Friendship
	err := r.db.Preload("User").Preload("Friend").
		Where("(user_id = ? OR friend_id = ?) AND status = ?", userID, userID, "accepted").
		Find(&friendships).Error
	return friendships, err
}

func (r *friendshipRepository) GetPendingRequests(userID uuid.UUID) ([]*entities.Friendship, error) {
	var friendships []*entities.Friendship
	err := r.db.Preload("User").Preload("Friend").
		Where("friend_id = ? AND status = ?", userID, "pending").
		Find(&friendships).Error
	return friendships, err
}

func (r *friendshipRepository) GetSentRequests(userID uuid.UUID) ([]*entities.Friendship, error) {
	var friendships []*entities.Friendship
	err := r.db.Preload("User").Preload("Friend").
		Where("user_id = ? AND status = ?", userID, "pending").
		Find(&friendships).Error
	return friendships, err
}

func (r *friendshipRepository) UpdateStatus(id uuid.UUID, status string) error {
	return r.db.Model(&entities.Friendship{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *friendshipRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Friendship{}, "id = ?", id).Error
}

func (r *friendshipRepository) AreFriends(userID, friendID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&entities.Friendship{}).
		Where("((user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)) AND status = ?",
			userID, friendID, friendID, userID, "accepted").
		Count(&count).Error
	return count > 0, err
}
