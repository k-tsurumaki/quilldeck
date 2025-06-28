package main

import (
	"fmt"
	"log"
	"os"

	"github/k-tsurumaki/quilldeck/internal/config"
	"github/k-tsurumaki/quilldeck/internal/infrastructure/database/sqlite"
	httpServer "github/k-tsurumaki/quilldeck/internal/interfaces/http"
)

func main() {
	cfg := config.Load()
	
	// データベース接続
	db, err := sqlite.NewConnection(cfg.Database.Path)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	
	// データディレクトリ作成
	if err := os.MkdirAll("./data", 0755); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}
	
	// マイグレーション実行
	if err := db.RunMigrations(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	
	server := httpServer.NewServer(db)
	
	fmt.Printf("Starting QuillDeck server on port %s...\n", cfg.Server.Port)
	
	if err := server.Start(cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}