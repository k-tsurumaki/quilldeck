package config

import (
	"os"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	LLM      LLMConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Type string
	Path string
}

type LLMConfig struct {
	LLM_API_KEY  string
	LLM_BASE_URL string
	LLM_MODEL   string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		Database: DatabaseConfig{
			Type: getEnv("DB_TYPE", "sqlite"),
			Path: getEnv("DB_PATH", "./data/quilldeck.db"),
		},
		LLM: LLMConfig{
			LLM_API_KEY:  getEnv("LLM_API_KEY", ""),
			LLM_BASE_URL: getEnv("LLM_BASE_URL", ""),
			LLM_MODEL: getEnv("LLM_MODEL", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
