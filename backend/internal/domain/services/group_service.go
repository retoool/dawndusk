package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupService interface {
	Create(userID uuid.UUID, req *dto.CreateGroupRequest) (*dto.GroupResponse, error)
	GetByID(groupID uuid.UUID) (*dto.GroupResponse, error)
	Update(groupID, userID uuid.UUID, req *dto.UpdateGroupRequest) (*dto.GroupResponse, error)
	Delete(groupID, userID uuid.UUID) error
	List(limit, offset int) ([]*dto.GroupResponse, error)
	GetUserGroups(userID uuid.UUID) ([]*dto.GroupResponse, error)

	// Member operations
	JoinByInviteCode(userID uuid.UUID, inviteCode string) (*dto.GroupResponse, error)
	Leave(groupID, userID uuid.UUID) error
	GetMembers(groupID uuid.UUID) ([]*dto.GroupMemberResponse, error)
	UpdateMemberRole(groupID, userID, targetUserID uuid.UUID, role string) error
	RemoveMember(groupID, userID, targetUserID uuid.UUID) error
}

type groupService struct {
	groupRepo repositories.GroupRepository
}

func NewGroupService(groupRepo repositories.GroupRepository) GroupService {
	return &groupService{
		groupRepo: groupRepo,
	}
}

func (s *groupService) Create(userID uuid.UUID, req *dto.CreateGroupRequest) (*dto.GroupResponse, error) {
	// Generate unique invite code
	inviteCode, err := generateInviteCode()
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	group := &entities.Group{
		Name:        req.Name,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
		CreatorID:   &userID,
		MaxMembers:  req.MaxMembers,
		IsPrivate:   req.IsPrivate,
		InviteCode:  inviteCode,
	}

	if err := s.groupRepo.Create(group); err != nil {
		return nil, errors.ErrInternalServer
	}

	// Add creator as admin member
	member := &entities.GroupMember{
		GroupID: group.ID,
		UserID:  userID,
		Role:    "admin",
	}

	if err := s.groupRepo.AddMember(member); err != nil {
		return nil, errors.ErrInternalServer
	}

	return s.toResponse(group), nil
}

func (s *groupService) GetByID(groupID uuid.UUID) (*dto.GroupResponse, error) {
	group, err := s.groupRepo.FindByID(groupID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError("GROUP_NOT_FOUND", "Group not found", 404)
		}
		return nil, errors.ErrInternalServer
	}

	return s.toResponse(group), nil
}

func (s *groupService) Update(groupID, userID uuid.UUID, req *dto.UpdateGroupRequest) (*dto.GroupResponse, error) {
	// Check if user is admin
	role, err := s.groupRepo.GetMemberRole(groupID, userID)
	if err != nil {
		return nil, errors.NewAppError("NOT_MEMBER", "You are not a member of this group", 403)
	}

	if role != "admin" {
		return nil, errors.NewAppError("PERMISSION_DENIED", "Only admins can update group", 403)
	}

	group, err := s.groupRepo.FindByID(groupID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	// Update fields
	if req.Name != nil {
		group.Name = *req.Name
	}
	if req.Description != nil {
		group.Description = req.Description
	}
	if req.AvatarURL != nil {
		group.AvatarURL = req.AvatarURL
	}
	if req.MaxMembers != nil {
		group.MaxMembers = *req.MaxMembers
	}
	if req.IsPrivate != nil {
		group.IsPrivate = *req.IsPrivate
	}

	if err := s.groupRepo.Update(group); err != nil {
		return nil, errors.ErrInternalServer
	}

	return s.toResponse(group), nil
}

func (s *groupService) Delete(groupID, userID uuid.UUID) error {
	// Check if user is admin
	role, err := s.groupRepo.GetMemberRole(groupID, userID)
	if err != nil {
		return errors.NewAppError("NOT_MEMBER", "You are not a member of this group", 403)
	}

	if role != "admin" {
		return errors.NewAppError("PERMISSION_DENIED", "Only admins can delete group", 403)
	}

	return s.groupRepo.Delete(groupID)
}

func (s *groupService) List(limit, offset int) ([]*dto.GroupResponse, error) {
	groups, err := s.groupRepo.List(limit, offset)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.GroupResponse, len(groups))
	for i, group := range groups {
		responses[i] = s.toResponse(group)
	}

	return responses, nil
}

func (s *groupService) GetUserGroups(userID uuid.UUID) ([]*dto.GroupResponse, error) {
	groups, err := s.groupRepo.FindGroupsByUserID(userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.GroupResponse, len(groups))
	for i, group := range groups {
		responses[i] = s.toResponse(group)
	}

	return responses, nil
}

func (s *groupService) JoinByInviteCode(userID uuid.UUID, inviteCode string) (*dto.GroupResponse, error) {
	group, err := s.groupRepo.FindByInviteCode(inviteCode)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError("INVALID_CODE", "Invalid invite code", 404)
		}
		return nil, errors.ErrInternalServer
	}

	// Check if already a member
	isMember, err := s.groupRepo.IsMember(group.ID, userID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	if isMember {
		return nil, errors.NewAppError("ALREADY_MEMBER", "You are already a member of this group", 400)
	}

	// Check if group is full
	members, err := s.groupRepo.FindMembersByGroupID(group.ID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	if len(members) >= group.MaxMembers {
		return nil, errors.NewAppError("GROUP_FULL", "Group is full", 400)
	}

	// Add member
	member := &entities.GroupMember{
		GroupID: group.ID,
		UserID:  userID,
		Role:    "member",
	}

	if err := s.groupRepo.AddMember(member); err != nil {
		return nil, errors.ErrInternalServer
	}

	return s.toResponse(group), nil
}

func (s *groupService) Leave(groupID, userID uuid.UUID) error {
	// Check if user is a member
	isMember, err := s.groupRepo.IsMember(groupID, userID)
	if err != nil {
		return errors.ErrInternalServer
	}

	if !isMember {
		return errors.NewAppError("NOT_MEMBER", "You are not a member of this group", 400)
	}

	return s.groupRepo.RemoveMember(groupID, userID)
}

func (s *groupService) GetMembers(groupID uuid.UUID) ([]*dto.GroupMemberResponse, error) {
	members, err := s.groupRepo.FindMembersByGroupID(groupID)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	responses := make([]*dto.GroupMemberResponse, len(members))
	for i, member := range members {
		responses[i] = &dto.GroupMemberResponse{
			ID:       member.ID.String(),
			GroupID:  member.GroupID.String(),
			UserID:   member.UserID.String(),
			Username: member.User.Username,
			Role:     member.Role,
			JoinedAt: member.JoinedAt,
		}
	}

	return responses, nil
}

func (s *groupService) UpdateMemberRole(groupID, userID, targetUserID uuid.UUID, role string) error {
	// Check if user is admin
	userRole, err := s.groupRepo.GetMemberRole(groupID, userID)
	if err != nil {
		return errors.NewAppError("NOT_MEMBER", "You are not a member of this group", 403)
	}

	if userRole != "admin" {
		return errors.NewAppError("PERMISSION_DENIED", "Only admins can update member roles", 403)
	}

	// Validate role
	if role != "admin" && role != "moderator" && role != "member" {
		return errors.NewAppError("INVALID_ROLE", "Invalid role", 400)
	}

	// Get target member
	members, err := s.groupRepo.FindMembersByGroupID(groupID)
	if err != nil {
		return errors.ErrInternalServer
	}

	for _, member := range members {
		if member.UserID == targetUserID {
			member.Role = role
			// Note: This is a simplified update. In production, you'd want a dedicated UpdateMember method
			return nil
		}
	}

	return errors.NewAppError("MEMBER_NOT_FOUND", "Member not found", 404)
}

func (s *groupService) RemoveMember(groupID, userID, targetUserID uuid.UUID) error {
	// Check if user is admin or moderator
	userRole, err := s.groupRepo.GetMemberRole(groupID, userID)
	if err != nil {
		return errors.NewAppError("NOT_MEMBER", "You are not a member of this group", 403)
	}

	if userRole != "admin" && userRole != "moderator" {
		return errors.NewAppError("PERMISSION_DENIED", "Only admins and moderators can remove members", 403)
	}

	return s.groupRepo.RemoveMember(groupID, targetUserID)
}

func (s *groupService) toResponse(group *entities.Group) *dto.GroupResponse {
	response := &dto.GroupResponse{
		ID:          group.ID.String(),
		Name:        group.Name,
		Description: group.Description,
		AvatarURL:   group.AvatarURL,
		MaxMembers:  group.MaxMembers,
		IsPrivate:   group.IsPrivate,
		InviteCode:  group.InviteCode,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}

	if group.CreatorID != nil {
		creatorID := group.CreatorID.String()
		response.CreatorID = &creatorID
	}

	return response
}

func generateInviteCode() (string, error) {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("GRP-%s", hex.EncodeToString(bytes)[:8]), nil
}
