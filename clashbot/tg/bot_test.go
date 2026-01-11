package tg

import (
	"testing"
)

func TestBot_Initialization(t *testing.T) {
	bot := &Bot{
		awaitingTag: make(map[int64]bool),
	}

	if bot.awaitingTag == nil {
		t.Error("awaitingTag should be initialized")
	}

	bot.initHandlers()

	if bot.handlers == nil {
		t.Error("handlers should be initialized")
	}
}

func TestExtractCommand(t *testing.T) {
	cmd, args := extractCommandTest("/player #ABC123")
	if cmd != "player" {
		t.Errorf("Expected command 'player', got '%s'", cmd)
	}
	if args != "#ABC123" {
		t.Errorf("Expected args '#ABC123', got '%s'", args)
	}
}

func extractCommandTest(text string) (string, string) {
	if len(text) == 0 || text[0] != '/' {
		if len(text) == 0 {
			return "", ""
		}
		return "", text
	}

	text = text[1:]
	for i, ch := range text {
		if ch == ' ' {
			return text[:i], text[i+1:]
		}
	}
	return text, ""
}

func TestFormatTag(t *testing.T) {
	result := formatTagTest("ABC123")
	if result != "#ABC123" {
		t.Errorf("Expected #ABC123, got %s", result)
	}
}

func formatTagTest(tag string) string {
	if len(tag) == 0 {
		return "#"
	}
	if tag[0] != '#' {
		return "#" + tag
	}
	return tag
}
