package main

import (
	"log"
	"os/user"

	"github.com/KornutaKM/news_aggregator_with_tg/internal/article"
	"github.com/KornutaKM/news_aggregator_with_tg/internal/config"
	"github.com/KornutaKM/news_aggregator_with_tg/internal/database"
	"github.com/KornutaKM/news_aggregator_with_tg/internal/subscription"
	"github.com/KornutaKM/news_aggregator_with_tg/internal/topic"
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

	db, err := database.Connect(cfg.DB.DSN)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		return
	}
	logger.Info("✅ Connected to PostgreSQL!")

	if err := db.AutoMigrate(
		&user.User{},
		&topic.Topic{},
		&article.Article{},
		&subscription.Subscription{},
	); err != nil {
		logger.Error("failed to migrate database", "error", err)
		return
	}
	logger.Info("✅ Database migrated!")

	logger.Info("🚀 News Aggregator 4 starting...")
}
