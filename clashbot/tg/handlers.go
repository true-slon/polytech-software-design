package tg

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) initHandlers() map[string]func(*tgbotapi.Message) {
	return map[string]func(*tgbotapi.Message){
		"start":     b.handleStart,
		"player":    b.handlePlayer,
		"clan":      b.handleClan,
		"battlelog": b.handleBattleLog,
		"admin":     b.handleAdmin,
	}
}

func (b *Bot) reply(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	b.api.Send(msg)
}

func (b *Bot) handleStart(msg *tgbotapi.Message) {
	b.addUserIfNotExists(msg.From)
	b.reply(msg.Chat.ID, "Пиши\n/player #TAG\n/clan #TAG\nзаебал")
}
func (b *Bot) handleAdmin(msg *tgbotapi.Message) {
	var res bool
	err := b.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM clashbot.Telegram_users WHERE id = $1 and is_admin = $2)",
		msg.From.ID,
		true,
	).Scan(&res)
	if err != nil {
		log.Printf("Error select: %v", err)

	}
	if res == true {
		b.reply(msg.Chat.ID, "да ты пиздец ахуевший")
	} else {
		b.reply(msg.Chat.ID, "пососи")
	}
}
func (b *Bot) addUserIfNotExists(user *tgbotapi.User) {
	_, err := b.db.Exec(
		`INSERT INTO clashbot.Telegram_users(id, first_name, username, date) VALUES($1, $2, $3, $4)`,
		user.ID,
		user.FirstName,
		user.UserName,
		time.Now().Unix(),
	)
	if err != nil {
		log.Printf("Error adding user: %v", err)
	}
}

func (b *Bot) handlePlayer(msg *tgbotapi.Message) {
	tag := msg.CommandArguments()
	if tag == "" {
		b.reply(msg.Chat.ID, "тег бро")
		return
	}

	player, err := b.service.GetPlayer(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	text := fmt.Sprintf(
		"Имя: %s\nКубки: %d\nАренa: %s\nЛюбимая карта: %s",
		player.Name,
		player.Trophies,
		player.Arena.Name,
		player.CurrentFavouriteCard.Name,
	)

	b.reply(msg.Chat.ID, text)
}

func (b *Bot) handleClan(msg *tgbotapi.Message) {
	tag := msg.CommandArguments()
	if tag == "" {
		b.reply(msg.Chat.ID, "тег бро")
		return
	}

	clan, err := b.service.GetClan(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	text := fmt.Sprintf(
		"Название: %s\nКубки: %d\nСтрана: %s\nОчко войны %d\nУчастников: %d\nОписание: %s",
		clan.Name,
		clan.ClanScore,
		clan.Location.Name,
		clan.ClanWarTrophies,
		clan.Members,
		clan.Description,
	)

	b.reply(msg.Chat.ID, text)
}

func (b *Bot) handleBattleLog(msg *tgbotapi.Message) {
	tag := msg.CommandArguments()
	if tag == "" {
		b.reply(msg.Chat.ID, "тег бро")
		return
	}

	battleLog, err := b.service.GetBattleLog(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	trophiesChange := 0
	for _, battle := range *battleLog {
		trophiesChange += battle.Team[0].TrophyChange
	}
	b.reply(msg.Chat.ID, "Бро... "+strconv.Itoa(trophiesChange))
}

func (b *Bot) handleUnknown(msg *tgbotapi.Message) {
	b.reply(msg.Chat.ID, "да")
}
