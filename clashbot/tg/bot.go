package tg

import (
	"database/sql"
	"log"

	"clashbot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	service  *service.Service
	db       *sql.DB
	handlers map[string]func(*tgbotapi.Message)
}

func NewBot(token string, service *service.Service, db *sql.DB) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		api:     api,
		service: service,
		db:      db,
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
func (b *Bot) initMenu(url string) error {
	// menu := &tgbotapi.MenuButton{
	// 	Type:   "web_app",
	// 	Text:   "Open",
	// 	WebApp: &tgbotapi.WebAppInfo{URL: url},
	// }

	// _, err := b.api.Request(tgbotapi.SetChatMenuButtonConfig{MenuButton: menu})

	// if err != nil {
	// 	return (err)
	// }
	// return nil
	return nil
}
