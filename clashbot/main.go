package main

import (
	"log"

	"clashbot/config"
	"clashbot/cr"
	"clashbot/service"
	"clashbot/tg"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load: %v", err)
	}

	crClient := cr.NewClient(cfg.ClashApiKey)
	svc := service.NewService(crClient)

	bot := tg.NewBot(cfg.BotToken, svc)

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
