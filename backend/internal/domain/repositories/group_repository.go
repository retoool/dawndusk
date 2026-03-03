package repositories

import (
	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupRepository interface {
	Create(group *entities.Group) error
	FindByID(id uuid.UUID) (*entities.Group, error)
	FindByInviteCode(code string) (*entities.Group, error)
	Update(group *entities.Group) error
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]*entities.Group, error)

	// Group member operations
	AddMember(member *entities.GroupMember) error
	RemoveMember(groupID, userID uuid.UUID) error
	FindMembersByGroupID(groupID uuid.UUID) ([]*entities.GroupMember, error)
	FindGroupsByUserID(userID uuid.UUID) ([]*entities.Group, error)
	IsMember(groupID, userID uuid.UUID) (bool, error)
	GetMemberRole(groupID, userID uuid.UUID) (string, error)
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) Create(group *entities.Group) error {
	return r.db.Create(group).Error
}

func (r *groupRepository) FindByID(id uuid.UUID) (*entities.Group, error) {
	var group entities.Group
	err := r.db.Preload("Creator").First(&group, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) FindByInviteCode(code string) (*entities.Group, error) {
	var group entities.Group
	err := r.db.First(&group, "invite_code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) Update(group *entities.Group) error {
	return r.db.Save(group).Error
}

func (r *groupRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Group{}, "id = ?", id).Error
}

func (r *groupRepository) List(limit, offset int) ([]*entities.Group, error) {
	var groups []*entities.Group
	err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&groups).Error
	return groups, err
}

func (r *groupRepository) AddMember(member *entities.GroupMember) error {
	return r.db.Create(member).Error
}

func (r *groupRepository) RemoveMember(groupID, userID uuid.UUID) error {
	return r.db.Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&entities.GroupMember{}).Error
}

func (r *groupRepository) FindMembersByGroupID(groupID uuid.UUID) ([]*entities.GroupMember, error) {
	var members []*entities.GroupMember
	err := r.db.Preload("User").Where("group_id = ?", groupID).Find(&members).Error
	return members, err
}

func (r *groupRepository) FindGroupsByUserID(userID uuid.UUID) ([]*entities.Group, error) {
	var groups []*entities.Group
	err := r.db.
		Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ?", userID).
		Find(&groups).Error
	return groups, err
}

func (r *groupRepository) IsMember(groupID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&entities.GroupMember{}).
		Where("group_id = ? AND user_id = ?", groupID, userID).
		Count(&count).Error
	return count > 0, err
}

func (r *groupRepository) GetMemberRole(groupID, userID uuid.UUID) (string, error) {
	var member entities.GroupMember
	err := r.db.Where("group_id = ? AND user_id = ?", groupID, userID).First(&member).Error
	if err != nil {
		return "", err
	}
	return member.Role, nil
}
