package main

import (
	"log"

	"github.com/KornutaKM/news_aggregator_with_tg/internal/config"
	"github.com/KornutaKM/news_aggregator_with_tg/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	logger, err := logger.New(cfg.Log)
	if err != nil {
		log.Fatal("failed to create logger:", err)
	}

	logger.Info("✅ Config loaded successfully")
	logger.Info("🚀 News Aggregator 4 starting...")
}
