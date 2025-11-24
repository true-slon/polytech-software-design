package tg

import (
	"log"
	"net/http"

	"clashbot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	service  *service.Service
	handlers map[string]func(*tgbotapi.Message)
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

func (b *Bot) Run(url string, port string) error {
	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// b.api.Debug = true

	// updates := b.api.GetUpdatesChan(u)

	// b.initHandlers()
	// if err := b.initMenu(url); err != nil {
	// 	return (err)
	// }

	b.api.Debug = true

	b.initHandlers()

	_, _ = b.api.Request(tgbotapi.DeleteWebhookConfig{DropPendingUpdates: true})
	updates := b.api.ListenForWebhook("/wh")

	if port == "" {
		port = "8080"
	}
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	wh, err := tgbotapi.NewWebhook(url + "/wh")

	if err != nil {
		return (err)
	}

	_, err = b.api.Request(wh)
	if err != nil {
		return (err)
	}

	if err := b.initMenu(url); err != nil {
		return (err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		cmd := update.Message.Command()

		if h, ok := b.handlers[cmd]; ok {
			h(update.Message)
		} else {
			b.handleUnknown(update.Message)
		}
	}

	return nil
}

func (b *Bot) initHandlers() {
	b.handlers = map[string]func(*tgbotapi.Message){
		"start":     b.handleStart,
		"player":    b.handlePlayer,
		"clan":      b.handleClan,
		"battlelog": b.handleBattleLog,
	}
}

func (b *Bot) initMenu(url string) error {
	menu := &tgbotapi.MenuButton{
		Type:   "web_app",
		Text:   "Open",
		WebApp: &tgbotapi.WebAppInfo{URL: url},
	}

	_, err := b.api.Request(tgbotapi.SetChatMenuButtonConfig{MenuButton: menu})

	if err != nil {
		return (err)
	}
	return nil
}
