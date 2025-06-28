package http

import (
	"net/http"

	"github.com/k-tsurumaki/fuselage"
	"github/k-tsurumaki/quilldeck/internal/domain/service"
	"github/k-tsurumaki/quilldeck/internal/infrastructure/database/sqlite"
	"github/k-tsurumaki/quilldeck/internal/interfaces/http/handlers"
)

type Server struct {
	router      *fuselage.Router
	authHandler *handlers.AuthHandler
	docHandler  *handlers.DocumentHandler
}

func NewServer(db *sqlite.DB) *Server {
	router := fuselage.New()
	
	// リポジトリ作成
	userRepo := sqlite.NewUserRepository(db)
	docRepo := sqlite.NewDocumentRepository(db)
	summaryRepo := sqlite.NewSummaryRepository(db)
	
	// サービス作成
	authService := service.NewAuthService(userRepo)
	docService := service.NewDocumentService(docRepo, summaryRepo)
	
	return &Server{
		router:      router,
		authHandler: handlers.NewAuthHandler(authService, docService),
		docHandler:  handlers.NewDocumentHandler(docService),
	}
}

func (s *Server) Start(port string) error {
	// ヘルスチェック
	s.router.GET("/health", s.healthHandler)
	
	// 認証エンドポイント
	s.router.POST("/api/auth/register", s.authHandler.Register)
	s.router.POST("/api/auth/login", s.authHandler.Login)
	
	// ドキュメントエンドポイント
	s.router.POST("/api/documents/upload", s.docHandler.Upload)
	s.router.POST("/api/documents/summary", s.docHandler.GenerateSummary)

	server := fuselage.NewServer(":"+port, s.router)
	return server.ListenAndServe()
}

func (s *Server) healthHandler(c *fuselage.Context) error {
	// CORSヘッダーを設定
	c.Response.Header().Set("Access-Control-Allow-Origin", "*")
	c.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "quilldeck",
	})
}

