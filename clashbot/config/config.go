package config

import (
	"errors"
	"os"
)

type Config struct {
	BotToken    string
	ClashApiKey string
	WebAppUrl   string
	Port        string
}

func Load() (*Config, error) {
	cfg := &Config{
		BotToken:    os.Getenv("BOT_TOKEN"),
		ClashApiKey: os.Getenv("API_KEY"),
		WebAppUrl:   os.Getenv("WEBAPP_URL"),
		Port:        os.Getenv("WH_PORT"),
	}

	if cfg.BotToken == "" {
		return nil, errors.New("Missing bot token")
	}
	if cfg.ClashApiKey == "" {
		return nil, errors.New("Missing clash key")
	}
	if cfg.WebAppUrl == "" {
		return nil, errors.New("Missing webapp url")
	}

	return cfg, nil
}
