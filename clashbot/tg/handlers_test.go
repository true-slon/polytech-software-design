package tg

import (
	"testing"
)

func TestBot_HandlersMap(t *testing.T) {
	bot := &Bot{}
	bot.initHandlers()

	expectedHandlers := []string{
		"start", "player", "clan", "battlelog", "cardstat", "settag",
	}

	for _, handler := range expectedHandlers {
		if _, exists := bot.handlers[handler]; !exists {
			t.Errorf("Handler '%s' should exist", handler)
		}
	}
}

func TestAwaitingTagManagement(t *testing.T) {
	bot := &Bot{
		awaitingTag: make(map[int64]bool),
	}

	chatID := int64(12345)

	bot.awaitingTag[chatID] = true
	if !bot.awaitingTag[chatID] {
		t.Error("Tag awaiting should be true")
	}

	delete(bot.awaitingTag, chatID)
	if bot.awaitingTag[chatID] {
		t.Error("Tag awaiting should be false")
	}
}
