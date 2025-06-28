package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(email, password, name string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) Validate() error {
	if u.Email == "" {
		return &ValidationError{Field: "email", Message: "email is required"}
	}
	if u.Password == "" {
		return &ValidationError{Field: "password", Message: "password is required"}
	}
	if u.Name == "" {
		return &ValidationError{Field: "name", Message: "name is required"}
	}
	return nil
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}