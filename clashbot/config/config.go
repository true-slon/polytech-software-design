package config

import (
	"errors"
	"os"
)

type Config struct {
	BotToken    string
	ClashApiKey string
}

func Load() (*Config, error) {
	cfg := &Config{
		BotToken:    os.Getenv("BOT_TOKEN"),
		ClashApiKey: os.Getenv("API_KEY"),
	}

	if cfg.BotToken == "" {
		return nil, errors.New("Missing bot token")
	}
	if cfg.ClashApiKey == "" {
		return nil, errors.New("Missing clash key")
	}

	return cfg, nil
}
