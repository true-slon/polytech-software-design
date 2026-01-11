package tg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) reply(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	b.api.Send(msg)
}

func (b *Bot) sendMainMenu(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("👤 Игрок", "player"),
			tgbotapi.NewInlineKeyboardButtonData("🏰 Клан", "clan"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⚔ Последние бои", "battlelog"),
			tgbotapi.NewInlineKeyboardButtonData("🃏 Статистика карт", "cardstat"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Выбери команду:")
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

func (b *Bot) handleStart(msg *tgbotapi.Message) {
	b.sendMainMenu(msg.Chat.ID)
}

func (b *Bot) handlePlayer(msg *tgbotapi.Message, tag string) {
	if tag == "" {
		b.reply(msg.Chat.ID, "Пожалуйста, введите тег игрока.")
		return
	}

	player, err := b.service.GetPlayer(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	text := fmt.Sprintf(
		"Имя: %s\nКубки: %d\nАренa: %s\nУровень: %d\nВсего игр: %d\nПобед: %d\nТри короны: %d\nПроцент побед: %.2f\nЛюбимая карта: %s\nКоличество пожертвованных карт: %d",
		player.Name,
		player.Trophies,
		player.Arena.Name,
		player.ExpLevel,
		player.BattleCount,
		player.Wins,
		player.ThreeCrownWins,
		float64(player.Wins)/float64(player.BattleCount)*100,
		player.CurrentFavouriteCard.Name,
		player.TotalDonations,
	)

	b.reply(msg.Chat.ID, text)

	imageUrl := player.CurrentFavouriteCard.IconUrls.Medium
	photoMsg := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileURL(imageUrl))
	if _, err := b.api.Send(photoMsg); err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	elixirCost := 0
	var cardNames []string
	for _, card := range player.CurrentDeck {
		cardNames = append(cardNames, card.Name)
		elixirCost += card.ElixirCost
	}

	deckText := "Текущая колода:\n"
	for _, name := range cardNames {
		deckText += name + "\n"
	}

	deckText += fmt.Sprintf("Средняя стоимость эликсира: %.2f", float64(elixirCost)/float64(len(player.CurrentDeck)))

	b.reply(msg.Chat.ID, deckText)
}

func (b *Bot) handleClan(msg *tgbotapi.Message, tag string) {
	if tag == "" {
		b.reply(msg.Chat.ID, "Пожалуйста, введите тег клана.")
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

func (b *Bot) handleBattleLog(msg *tgbotapi.Message, tag string) {
	if tag == "" {
		b.reply(msg.Chat.ID, "Пожалуйста, введите тег игрока.")
		return
	}

	battleLog, err := b.service.GetBattleLog(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	wins := 0
	losses := 0
	for _, battle := range *battleLog {
		if battle.Team[0].TrophyChange > 0 {
			wins++
		} else {
			losses++
		}
	}

	if wins+losses == 0 {
		b.reply(msg.Chat.ID, "Слишком мало боёв для анализа!")
		return
	}

	trophiesChange := 0
	for _, battle := range *battleLog {
		trophiesChange += battle.Team[0].TrophyChange
	}

	text := fmt.Sprintf(
		"Статистика последних боёв:\n\nВсего боёв: %d\nПобед: %d\nПоражений: %d\nПроцент побед: %.2f\nИзменение кубков: %d",
		wins+losses,
		wins,
		losses,
		float64(wins)/float64(wins+losses)*100,
		trophiesChange,
	)

	b.reply(msg.Chat.ID, text)
}

func (b *Bot) handleCardStats(msg *tgbotapi.Message, tag string) {
	if tag == "" {
		b.reply(msg.Chat.ID, "Пожалуйста, введите тег игрока.")
		return
	}

	cardsStat, err := b.service.GetCardStats(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	text := fmt.Sprintf(
		"Анализ карт за последние бои:\n\nХудший противник: %s (WR:%.2f)\nЧасто встречается: %s (WR:%.2f)",
		cardsStat.WorstCard.Card.Name,
		cardsStat.WorstCard.Winrate*100,
		cardsStat.FrequentCard.Card.Name,
		cardsStat.FrequentCard.Winrate*100,
	)

	b.reply(msg.Chat.ID, text)

	imageUrl := cardsStat.WorstCard.Card.IconUrls.Medium
	photoMsg := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileURL(imageUrl))
	if _, err := b.api.Send(photoMsg); err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

}

func (b *Bot) handleUnknown(msg *tgbotapi.Message) {
	b.reply(msg.Chat.ID, "Неизвестная команда!")
}
