package config

import (
	"errors"
	"os"
)

type Config struct {
	BotToken         string
	ClashApiKey      string
	DatabasePassword string
	DatabaseUser     string
	DatabaseName     string
	DatabasePort     string
	WebAppUrl        string
}

func Load() (*Config, error) {
	cfg := &Config{
		BotToken:         os.Getenv("BOT_TOKEN"),
		ClashApiKey:      os.Getenv("CLASH_API_KEY"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
		DatabaseUser:     os.Getenv("DB_USER"),
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabasePort:     os.Getenv("DB_EXTERNAL_PORT"),
		WebAppUrl:        os.Getenv("WEBAPP_URL"),
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
