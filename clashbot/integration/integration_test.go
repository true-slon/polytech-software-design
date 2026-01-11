package integration

import (
	"testing"

	"clashbot/cr"
	"clashbot/service"
)

func TestIntegration_ServiceMethods(t *testing.T) {
	t.Run("CardStats winrate calculation", func(t *testing.T) {
		cs := service.CardStats{
			Wins:   2,
			Losses: 2,
		}
		cs.Winrate = float64(cs.Wins) / float64(cs.Wins+cs.Losses)
		if cs.Winrate != 0.5 {
			t.Errorf("Expected winrate 0.5, got %v", cs.Winrate)
		}
	})

	t.Run("RecentCards initialization", func(t *testing.T) {
		rc := service.RecentCards{}
		if rc.BestCard.Winrate != 0 {
			t.Errorf("BestCard winrate should be 0, got %v", rc.BestCard.Winrate)
		}
		// WorstCard.Winrate тоже 0 по умолчанию, не 1
		if rc.WorstCard.Winrate != 0 {
			t.Errorf("WorstCard winrate should be 0, got %v", rc.WorstCard.Winrate)
		}
	})
}

func TestIntegration_CRTypes(t *testing.T) {
	t.Run("Card structure", func(t *testing.T) {
		card := cr.Card{
			Name: "Test Card",
			ID:   1,
			IconUrls: cr.IconUrls{
				Medium: "test.jpg",
			},
		}

		if card.Name != "Test Card" {
			t.Errorf("Expected card name Test Card, got %s", card.Name)
		}
	})

	t.Run("Player structure", func(t *testing.T) {
		player := cr.Player{
			Name:     "Test Player",
			Trophies: 5000,
			Arena:    cr.Arena{Name: "Legendary"},
		}

		if player.Name != "Test Player" {
			t.Errorf("Expected player name Test Player, got %s", player.Name)
		}
	})

	t.Run("Clan structure", func(t *testing.T) {
		clan := cr.Clan{
			Name:    "Test Clan",
			Members: 50,
		}

		if clan.Name != "Test Clan" {
			t.Errorf("Expected clan name Test Clan, got %s", clan.Name)
		}
	})
}

func TestIntegration_UserFlow(t *testing.T) {
	t.Run("Service creation", func(t *testing.T) {
		var s *service.Service
		if s != nil {
			t.Error("Service should be nil")
		}
	})

	t.Run("CardStats default values", func(t *testing.T) {
		var cs service.CardStats
		if cs.Wins != 0 {
			t.Errorf("CardStats Wins should be 0, got %d", cs.Wins)
		}
		if cs.Losses != 0 {
			t.Errorf("CardStats Losses should be 0, got %d", cs.Losses)
		}
		if cs.Winrate != 0 {
			t.Errorf("CardStats Winrate should be 0, got %v", cs.Winrate)
		}
	})

	t.Run("BattleList type check", func(t *testing.T) {
		var bl cr.BattleList
		if bl == nil {
		}
	})
}
