package tg

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) reply(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	b.api.Send(msg)
}

func (b *Bot) handleStart(msg *tgbotapi.Message) {
	b.reply(msg.Chat.ID, "Пиши\n/player #TAG\n/clan #TAG\nзаебал")
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

func (b *Bot) handleCardStats(msg *tgbotapi.Message) {
	tag := msg.CommandArguments()
	if tag == "" {
		b.reply(msg.Chat.ID, "тег бро")
		return
	}

	cardsStat, err := b.service.GetCardStats(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	text := fmt.Sprintf(
		"Худший противник: %s(WR:%f)\nЧасто встречается: %s(WR:%f)",
		cardsStat.WorstCard.Card.Name,
		cardsStat.WorstCard.Winrate*100,
		cardsStat.FrequentCard.Card.Name,
		cardsStat.FrequentCard.Winrate*100,
	)

	b.reply(msg.Chat.ID, text)
}

func (b *Bot) handleUnknown(msg *tgbotapi.Message) {
	b.reply(msg.Chat.ID, "Иди нахуй")
}
