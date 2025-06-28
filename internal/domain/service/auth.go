package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github/k-tsurumaki/quilldeck/internal/domain/repository"
	"github/k-tsurumaki/quilldeck/internal/pkg/errors"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(ctx context.Context, email, password, name string) (*models.User, error) {
	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, errors.New(errors.ErrCodeValidation, "email already exists")
	}

	hashedPassword := s.hashPassword(password)
	user := models.NewUser(email, hashedPassword, name)

	if err := user.Validate(); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeValidation, "invalid user data")
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to create user")
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New(errors.ErrCodeUnauthorized, "invalid credentials")
	}

	hashedPassword := s.hashPassword(password)
	if user.Password != hashedPassword {
		return nil, errors.New(errors.ErrCodeUnauthorized, "invalid credentials")
	}

	return user, nil
}

func (s *AuthService) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeNotFound, "user not found")
	}
	return user, nil
}

func (s *AuthService) hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}