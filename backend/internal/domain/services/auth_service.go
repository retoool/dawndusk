package services

import (
	"time"

	"github.com/dawndusk/backend/internal/domain/entities"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/dto"
	"github.com/dawndusk/backend/internal/shared/config"
	"github.com/dawndusk/backend/internal/shared/errors"
	"github.com/dawndusk/backend/internal/shared/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
	RefreshToken(refreshToken string) (*dto.AuthResponse, error)
}

type authService struct {
	userRepo repositories.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo repositories.UserRepository, cfg *config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.ErrInternalServer
	}
	if existingUser != nil {
		return nil, errors.ErrEmailAlreadyExists
	}

	// Check if username already exists
	existingUser, err = s.userRepo.FindByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.ErrInternalServer
	}
	if existingUser != nil {
		return nil, errors.NewAppError("USERNAME_EXISTS", "Username already exists", 400)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	// Create user
	user := &entities.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Timezone:     "UTC",
		IsActive:     true,
		IsVerified:   false,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.ErrInternalServer
	}

	// Generate tokens
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, s.cfg)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, s.cfg)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.UserDTO{
			ID:         user.ID.String(),
			Username:   user.Username,
			Email:      user.Email,
			AvatarURL:  user.AvatarURL,
			Timezone:   user.Timezone,
			IsVerified: user.IsVerified,
		},
	}, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrInvalidCredentials
		}
		return nil, errors.ErrInternalServer
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, errors.ErrInvalidCredentials
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.NewAppError("USER_INACTIVE", "User account is inactive", 403)
	}

	// Update last login time
	now := time.Now()
	user.LastLoginAt = &now
	if err := s.userRepo.Update(user); err != nil {
		// Log error but don't fail login
	}

	// Generate tokens
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, s.cfg)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, s.cfg)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.UserDTO{
			ID:         user.ID.String(),
			Username:   user.Username,
			Email:      user.Email,
			AvatarURL:  user.AvatarURL,
			Timezone:   user.Timezone,
			IsVerified: user.IsVerified,
		},
	}, nil
}

func (s *authService) RefreshToken(refreshToken string) (*dto.AuthResponse, error) {
	// Validate refresh token
	claims, err := utils.ValidateToken(refreshToken, s.cfg.JWT.RefreshTokenSecret)
	if err != nil {
		return nil, errors.ErrUnauthorized
	}

	// Parse user ID
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, errors.ErrUnauthorized
	}

	// Find user
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrInternalServer
	}

	// Generate new tokens
	newAccessToken, err := utils.GenerateAccessToken(user.ID, user.Email, s.cfg)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	newRefreshToken, err := utils.GenerateRefreshToken(user.ID, s.cfg)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return &dto.AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		User: dto.UserDTO{
			ID:         user.ID.String(),
			Username:   user.Username,
			Email:      user.Email,
			AvatarURL:  user.AvatarURL,
			Timezone:   user.Timezone,
			IsVerified: user.IsVerified,
		},
	}, nil
}
