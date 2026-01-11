package tg

import (
	"database/sql"
	"log"

	"clashbot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api         *tgbotapi.BotAPI
	service     *service.Service
	db          *sql.DB
	handlers    map[string]func(*tgbotapi.Message)
	awaitingTag map[int64]bool
}

func NewBot(token string, service *service.Service, db *sql.DB) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot := &Bot{
		api:         api,
		service:     service,
		db:          db,
		awaitingTag: make(map[int64]bool),
	}

	bot.initHandlers()
	return bot
}

func (b *Bot) initHandlers() {
	b.handlers = map[string]func(*tgbotapi.Message){
		"start":     b.handleStart,
		"player":    b.handlePlayer,
		"clan":      b.handleClan,
		"battlelog": b.handleBattleLog,
		"cardstat":  b.handleCardStats,
		"settag":    b.handleSetTag,
	}
}

func (b *Bot) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)
	b.api.Debug = true

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if b.awaitingTag[update.Message.Chat.ID] {
			b.processTagInput(update.Message)
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
