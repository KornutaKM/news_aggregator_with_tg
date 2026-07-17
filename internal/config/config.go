package config

import (
	"fmt"
	"os"
	"time"

	"github.com/KornutaKM/news_aggregator_with_tg/pkg/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	DB       DBConfig
	JWT      JWTConfig
	Redis    RedisConfig
	Telegram TelegramConfig
	Log      logger.Config
}

type DBConfig struct {
	DSN string
}

type JWTConfig struct {
	Secret string
	TTL    time.Duration
}

type RedisConfig struct {
	Addr string
}

type TelegramConfig struct {
	Token string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	ttl, err := time.ParseDuration(os.Getenv("JWT_TTL"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse TTL: %w", err)
	}

	return &Config{
		DB: DBConfig{
			DSN: os.Getenv("DSN"),
		},
		Redis: RedisConfig{
			Addr: os.Getenv("REDIS_ADDR"),
		},
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
			TTL:    ttl,
		},
		Telegram: TelegramConfig{
			Token: os.Getenv("TELEGRAM_TOKEN"),
		},
		Log: logger.Config{
			Level:      os.Getenv("LOG_LEVEL"),
			FilePath:   os.Getenv("LOG_FILE"),
			MaxSize:    50,
			MaxBackups: 3,
			MaxAge:     30,
			Compress:   true,
			IsDev:      os.Getenv("ENV") != "production",
		},
	}, nil
}
