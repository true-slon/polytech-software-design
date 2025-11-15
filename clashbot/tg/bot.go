package tg

import (
	"log"

	"clashbot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api     *tgbotapi.BotAPI
	service *service.Service
}

func NewBot(token string, service *service.Service) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		api:     api,
		service: service,
	}
}

func (b *Bot) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	handlers := b.initHandlers()

	b.api.Debug = true

	for update := range updates {
		if update.Message == nil {
			continue
		}

		cmd := update.Message.Command()

		if h, ok := handlers[cmd]; ok {
			h(update.Message)
		} else {
			b.handleUnknown(update.Message)
		}
	}

	return nil
}
