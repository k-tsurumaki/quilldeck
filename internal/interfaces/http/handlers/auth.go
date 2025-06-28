package handlers

import (
	"net/http"

	"github.com/k-tsurumaki/fuselage"
	"github/k-tsurumaki/quilldeck/internal/domain/service"
)

type AuthHandler struct {
	authService *service.AuthService
	docService  *service.DocumentService
}

func NewAuthHandler(authService *service.AuthService, docService *service.DocumentService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		docService:  docService,
	}
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

	user, err := h.authService.Register(c.Request.Context(), req.Email, req.Password, req.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, AuthResponse{
		Message: "User registered successfully",
		UserID:  user.ID.String(),
	})
}

func (h *AuthHandler) Login(c *fuselage.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	user, err := h.authService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, AuthResponse{
		Message: "Login successful",
		UserID:  user.ID.String(),
	})
}