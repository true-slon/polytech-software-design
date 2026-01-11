package tg

import (
	"log"
	"net/http"
	"strings"

	"clashbot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	service  *service.Service
	handlers map[string]func(*tgbotapi.Message, string)

	waiting map[int64]string
}

func NewBot(token string, service *service.Service) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		api:     api,
		service: service,
		waiting: make(map[int64]string),
	}
}

func (b *Bot) Run(url string, port string) error {
	b.api.Debug = true

	b.initHandlers()

	_, _ = b.api.Request(tgbotapi.DeleteWebhookConfig{DropPendingUpdates: true})
	updates := b.api.ListenForWebhook("/wh")

	if port == "" {
		port = "8080"
	}
	go func() error {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			return err
		}
		return nil
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

		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			cmd := update.CallbackQuery.Data

			b.waiting[chatID] = cmd

			b.api.Send(tgbotapi.NewMessage(chatID, "Введи тег игрока / клана:"))

			b.api.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			continue
		}

		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		if cmd, ok := b.waiting[chatID]; ok {
			delete(b.waiting, chatID)

			tag := strings.TrimSpace(update.Message.Text)

			if h, ok := b.handlers[cmd]; ok {
				h(update.Message, tag)
			}
			continue
		}

		b.sendMainMenu(chatID)
	}

	return nil
}

func (b *Bot) initHandlers() {
	b.handlers = map[string]func(*tgbotapi.Message, string){
		"player":    b.handlePlayer,
		"clan":      b.handleClan,
		"battlelog": b.handleBattleLog,
		"cardstat":  b.handleCardStats,
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
