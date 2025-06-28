package http

import (
	"net/http"

	"github.com/k-tsurumaki/fuselage"
	"github/k-tsurumaki/quilldeck/internal/interfaces/http/handlers"
)

type Server struct {
	router      *fuselage.Router
	authHandler *handlers.AuthHandler
}

func NewServer() *Server {
	router := fuselage.New()
	
	return &Server{
		router:      router,
		authHandler: handlers.NewAuthHandler(),
	}
}

func (s *Server) Start(port string) error {
	// ヘルスチェック
	s.router.GET("/health", s.healthHandler)
	
	// 認証エンドポイント
	s.router.POST("/api/auth/register", s.authHandler.Register)
	s.router.POST("/api/auth/login", s.authHandler.Login)

	server := fuselage.NewServer(":"+port, s.router)
	return server.ListenAndServe()
}

func (s *Server) healthHandler(c *fuselage.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "quilldeck",
	})
}