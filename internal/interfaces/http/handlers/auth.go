package handlers

import (
	"net/http"

	"github.com/k-tsurumaki/fuselage"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Message string `json:"message"`
	UserID  string `json:"user_id,omitempty"`
}

func (h *AuthHandler) Register(c *fuselage.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	return c.JSON(http.StatusOK, AuthResponse{
		Message: "Registration endpoint working",
		UserID:  "test-user-id",
	})
}

func (h *AuthHandler) Login(c *fuselage.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	return c.JSON(http.StatusOK, AuthResponse{
		Message: "Login endpoint working",
		UserID:  "test-user-id",
	})
}