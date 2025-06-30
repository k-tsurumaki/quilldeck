package http

import (
	"net/http"
	"github/k-tsurumaki/quilldeck/internal/config"

	"github.com/k-tsurumaki/fuselage"
	"github.com/k-tsurumaki/fuselage/middleware"
	"github/k-tsurumaki/quilldeck/internal/domain/service"
	"github/k-tsurumaki/quilldeck/internal/infrastructure/database/sqlite"
	"github/k-tsurumaki/quilldeck/internal/interfaces/http/handlers"
)

type Server struct {
	router      *fuselage.Router
	authHandler *handlers.AuthHandler
	docHandler  *handlers.DocumentHandler
}

func NewServer(db *sqlite.DB, cfg *config.Config) *Server {
	router := fuselage.New()
	
	// Add CORS middleware
	router.Use(middleware.CORS())
	
	// Create repositories
	userRepo := sqlite.NewUserRepository(db)
	docRepo := sqlite.NewDocumentRepository(db)
	summaryRepo := sqlite.NewSummaryRepository(db)
	
	// Create services
	authService := service.NewAuthService(userRepo)
	docService := service.NewDocumentService(docRepo, summaryRepo, cfg.LLM.LLM_API_KEY, cfg.LLM.LLM_BASE_URL, cfg.LLM.LLM_MODEL)
	
	return &Server{
		router:      router,
		authHandler: handlers.NewAuthHandler(authService, docService),
		docHandler:  handlers.NewDocumentHandler(docService),
	}
}

func (s *Server) Start(port string) error {
	// Health check endpoint
	s.router.GET("/health", s.healthHandler)
	
	// Authentication endpoints
	s.router.POST("/api/auth/register", s.authHandler.Register)
	s.router.POST("/api/auth/login", s.authHandler.Login)
	
	// Document endpoints
	s.router.POST("/api/documents/upload", s.docHandler.Upload)
	s.router.POST("/api/documents/summary", s.docHandler.GenerateSummary)

	server := fuselage.NewServer(":"+port, s.router)
	return server.ListenAndServe()
}

func (s *Server) healthHandler(c *fuselage.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "quilldeck",
	})
}

