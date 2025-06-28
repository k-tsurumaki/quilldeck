package service

import (
	"context"
	"errors"
	"testing"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github/k-tsurumaki/quilldeck/internal/domain/service"
	"github/k-tsurumaki/quilldeck/tests/unit/domain/service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Register(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string
		userName string
		setup    func(*mocks.MockUserRepository)
		wantErr  bool
		errCode  string
	}{
		{
			name:     "successful registration",
			email:    "test@example.com",
			password: "password123",
			userName: "Test User",
			setup: func(repo *mocks.MockUserRepository) {
				repo.On("GetByEmail", mock.Anything, "test@example.com").Return(nil, errors.New("not found"))
				repo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)
			},
			wantErr: false,
		},
		{
			name:     "email already exists",
			email:    "existing@example.com",
			password: "password123",
			userName: "Test User",
			setup: func(repo *mocks.MockUserRepository) {
				existingUser := models.NewUser("existing@example.com", "hashedpass", "Existing User")
				repo.On("GetByEmail", mock.Anything, "existing@example.com").Return(existingUser, nil)
			},
			wantErr: true,
			errCode: "VALIDATION_ERROR",
		},
		{
			name:     "invalid user data - empty email",
			email:    "",
			password: "password123",
			userName: "Test User",
			setup: func(repo *mocks.MockUserRepository) {
				repo.On("GetByEmail", mock.Anything, "").Return(nil, errors.New("not found"))
			},
			wantErr: true,
			errCode: "VALIDATION_ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.MockUserRepository)
			tt.setup(mockRepo)

			authService := service.NewAuthService(mockRepo)
			user, err := authService.Register(context.Background(), tt.email, tt.password, tt.userName)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
				if tt.errCode != "" {
					assert.Contains(t, err.Error(), tt.errCode)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.email, user.Email)
				assert.Equal(t, tt.userName, user.Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	hashedPassword := "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f" // sha256 of "password123"

	tests := []struct {
		name     string
		email    string
		password string
		setup    func(*mocks.MockUserRepository)
		wantErr  bool
		errCode  string
	}{
		{
			name:     "successful login",
			email:    "test@example.com",
			password: "password123",
			setup: func(repo *mocks.MockUserRepository) {
				user := &models.User{
					ID:       uuid.New(),
					Email:    "test@example.com",
					Password: hashedPassword,
					Name:     "Test User",
				}
				repo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
			},
			wantErr: false,
		},
		{
			name:     "user not found",
			email:    "notfound@example.com",
			password: "password123",
			setup: func(repo *mocks.MockUserRepository) {
				repo.On("GetByEmail", mock.Anything, "notfound@example.com").Return(nil, errors.New("not found"))
			},
			wantErr: true,
			errCode: "UNAUTHORIZED",
		},
		{
			name:     "wrong password",
			email:    "test@example.com",
			password: "wrongpassword",
			setup: func(repo *mocks.MockUserRepository) {
				user := &models.User{
					ID:       uuid.New(),
					Email:    "test@example.com",
					Password: hashedPassword,
					Name:     "Test User",
				}
				repo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
			},
			wantErr: true,
			errCode: "UNAUTHORIZED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.MockUserRepository)
			tt.setup(mockRepo)

			authService := service.NewAuthService(mockRepo)
			user, err := authService.Login(context.Background(), tt.email, tt.password)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
				if tt.errCode != "" {
					assert.Contains(t, err.Error(), tt.errCode)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.email, user.Email)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}