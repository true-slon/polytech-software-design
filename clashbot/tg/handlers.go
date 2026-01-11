package tg

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) reply(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	b.api.Send(msg)
}

func (b *Bot) saveUser(userID int64, firstName, username string, date int64) error {
	query := `
		INSERT INTO clashbot.telegram_users (id, first_name, username, date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		ON CONFLICT (id) DO UPDATE SET
			first_name = $2,
			username = $3,
			updated_at = NOW()
	`
	_, err := b.db.Exec(query, userID, firstName, username, date)
	return err
}

func (b *Bot) saveClashTag(userID int64, tag string) error {
	cleanedTag := strings.TrimPrefix(tag, "#")

	query := `
		UPDATE clashbot.telegram_users 
		SET clash_tag = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := b.db.Exec(query, cleanedTag, userID)
	return err
}

func (b *Bot) getClashTag(userID int64) (string, error) {
	var tag string
	query := `SELECT clash_tag FROM clashbot.telegram_users WHERE id = $1`
	err := b.db.QueryRow(query, userID).Scan(&tag)

	if err == sql.ErrNoRows {
		return "", nil
	}

	return tag, err
}

func (b *Bot) handleStart(msg *tgbotapi.Message) {
	err := b.saveUser(msg.From.ID, msg.From.FirstName, msg.From.UserName, int64(msg.Date))
	if err != nil {
		log.Printf("Error saving user: %v", err)
	}

	tag, err := b.getClashTag(msg.From.ID)
	if err != nil {
		log.Printf("Error getting tag: %v", err)
	}

	if tag == "" {
		b.reply(msg.Chat.ID, "Привет! Для начала отправь мне свой тег Clash Royale (например, #ABC123).")
		b.awaitingTag[msg.Chat.ID] = true
	} else {
		b.reply(msg.Chat.ID, "Привет! Твой сохраненный тег: #"+tag+"\n\nКоманды:\n/player - информация об игроке\n/clan #ТЕГ - информация о клане (укажи тег клана)\n/battlelog - статистика боёв\n/cardstat - статистика карт\n/settag - изменить тег")
	}
}

func (b *Bot) processTagInput(msg *tgbotapi.Message) {
	tag := strings.TrimSpace(msg.Text)

	if tag == "" {
		b.reply(msg.Chat.ID, "Тег не может быть пустым. Отправь мне свой тег Clash Royale (например, #ABC123).")
		return
	}

	delete(b.awaitingTag, msg.Chat.ID)

	if !strings.HasPrefix(tag, "#") {
		tag = "#" + tag
	}

	err := b.saveClashTag(msg.From.ID, tag)
	if err != nil {
		b.reply(msg.Chat.ID, "Ошибка сохранения тега. Попробуй еще раз.")
		log.Printf("Error saving tag: %v", err)
		b.awaitingTag[msg.Chat.ID] = true
		return
	}

	b.reply(msg.Chat.ID, "Тег "+tag+" сохранён!\n\nКоманды:\n/player - информация об игроке\n/clan #ТЕГ - информация о клане (укажи тег клана)\n/battlelog - статистика боёв\n/cardstat - статистика карт\n/settag - изменить тег")
}

func (b *Bot) handleSetTag(msg *tgbotapi.Message) {
	b.reply(msg.Chat.ID, "Отправь мне новый тег Clash Royale (например, #ABC123).")
	b.awaitingTag[msg.Chat.ID] = true
}

func (b *Bot) handlePlayer(msg *tgbotapi.Message) {
	tag := msg.CommandArguments()

	if tag == "" {
		savedTag, err := b.getClashTag(msg.From.ID)
		if err != nil || savedTag == "" {
			b.reply(msg.Chat.ID, "У тебя нет сохраненного тега. Используй /start или укажи тег: /player #TAG")
			return
		}
		tag = savedTag
	} else if !strings.HasPrefix(tag, "#") {
		tag = "#" + tag
	}

	player, err := b.service.GetPlayer(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	text := fmt.Sprintf(
		"Имя: %s\nКубки: %d\nАрена: %s\nЛюбимая карта: %s",
		player.Name,
		player.Trophies,
		player.Arena.Name,
		player.CurrentFavouriteCard.Name,
	)

	b.reply(msg.Chat.ID, text)

	imageUrl := player.CurrentFavouriteCard.IconUrls.Medium
	photoMsg := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileURL(imageUrl))
	b.api.Send(photoMsg)

	var cardNames []string
	for _, card := range player.CurrentDeck {
		cardNames = append(cardNames, card.Name)
	}

	deckText := "Текущая колода:\n" + fmt.Sprintf("%s", cardNames)
	b.reply(msg.Chat.ID, deckText)
}

func (b *Bot) handleClan(msg *tgbotapi.Message) {
	tag := msg.CommandArguments()

	// ВАЖНО: Убираем проверку на сохраненный тег!
	// Команда /clan всегда требует явного указания тега клана
	if tag == "" {
		b.reply(msg.Chat.ID, "Укажи тег клана для поиска: /clan #ТЕГ_КЛАНА\n\nТы можешь найти тег клана в игре в описании клана.")
		return
	}

	if !strings.HasPrefix(tag, "#") {
		tag = "#" + tag
	}

	clan, err := b.service.GetClan(tag)
	if err != nil {
		b.reply(msg.Chat.ID, "Не удалось найти клан. Проверь правильность тега.")
		log.Printf("Error getting clan for tag %s: %v", tag, err)
		return
	}

	text := fmt.Sprintf(
		"Название: %s\nКубки: %d\nСтрана: %s\nОчко войны: %d\nУчастников: %d\nОписание: %s",
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
		savedTag, err := b.getClashTag(msg.From.ID)
		if err != nil || savedTag == "" {
			b.reply(msg.Chat.ID, "У тебя нет сохраненного тега. Используй /start или укажи тег: /battlelog #TAG")
			return
		}
		tag = savedTag
	} else if !strings.HasPrefix(tag, "#") {
		tag = "#" + tag
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
		"Всего боёв: %d\nПобед: %d\nПоражений: %d\nПроцент побед: %.1f%%\nИзменение кубков: %d",
		wins+losses,
		wins,
		losses,
		float64(wins)/float64(wins+losses)*100,
		trophiesChange,
	)

	b.reply(msg.Chat.ID, text)
}

func (b *Bot) handleCardStats(msg *tgbotapi.Message) {
	tag := msg.CommandArguments()

	if tag == "" {
		savedTag, err := b.getClashTag(msg.From.ID)
		if err != nil || savedTag == "" {
			b.reply(msg.Chat.ID, "У тебя нет сохраненного тега. Используй /start или укажи тег: /cardstat #TAG")
			return
		}
		tag = savedTag
	} else if !strings.HasPrefix(tag, "#") {
		tag = "#" + tag
	}

	cardsStat, err := b.service.GetCardStats(tag)
	if err != nil {
		b.reply(msg.Chat.ID, err.Error())
		return
	}

	text := fmt.Sprintf(
		"Худший противник: %s(WR:%.1f%%)\nЧасто встречается: %s(WR:%.1f%%)",
		cardsStat.WorstCard.Card.Name,
		cardsStat.WorstCard.Winrate*100,
		cardsStat.FrequentCard.Card.Name,
		cardsStat.FrequentCard.Winrate*100,
	)

	b.reply(msg.Chat.ID, text)

	imageUrl := cardsStat.WorstCard.Card.IconUrls.Medium
	photoMsg := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileURL(imageUrl))
	b.api.Send(photoMsg)
}

func (b *Bot) handleUnknown(msg *tgbotapi.Message) {
	b.reply(msg.Chat.ID, "Неизвестная команда. Используй /start")
}
