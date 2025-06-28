package models

import (
	"testing"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	email := "test@example.com"
	password := "password123"
	name := "Test User"

	user := models.NewUser(email, password, name)

	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, name, user.Name)
	assert.NotEqual(t, "", user.ID.String())
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid user",
			user: models.NewUser("test@example.com", "password123", "Test User"),
			wantErr: false,
		},
		{
			name: "empty email",
			user: &models.User{
				Email:    "",
				Password: "password123",
				Name:     "Test User",
			},
			wantErr: true,
			errMsg:  "email is required",
		},
		{
			name: "empty password",
			user: &models.User{
				Email:    "test@example.com",
				Password: "",
				Name:     "Test User",
			},
			wantErr: true,
			errMsg:  "password is required",
		},
		{
			name: "empty name",
			user: &models.User{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "",
			},
			wantErr: true,
			errMsg:  "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}