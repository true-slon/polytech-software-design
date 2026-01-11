package service

import (
	"clashbot/cr"
	"testing"
)

type MockClient struct{}

func (m *MockClient) GetPlayer(tag string) (*cr.Player, error) {
	return &cr.Player{
		Name:                 "TestPlayer",
		Trophies:             5000,
		Arena:                cr.Arena{Name: "Test Arena"},
		CurrentFavouriteCard: &cr.Card{Name: "Test Card"},
	}, nil
}

func (m *MockClient) GetClan(tag string) (*cr.Clan, error) {
	return &cr.Clan{
		Name:        "Test Clan",
		Members:     50,
		Description: "Test Description",
	}, nil
}

func (m *MockClient) GetBattleLog(tag string) (*cr.BattleList, error) {
	return &cr.BattleList{}, nil
}

func TestCardStats_CalculateWinrate(t *testing.T) {
	cs := CardStats{
		Wins:   3,
		Losses: 1,
	}
	cs.Winrate = float64(cs.Wins) / float64(cs.Wins+cs.Losses)
	if cs.Winrate != 0.75 {
		t.Errorf("Expected winrate 0.75, got %v", cs.Winrate)
	}
}

func TestCardStats_NoGames(t *testing.T) {
	cs := CardStats{
		Wins:   0,
		Losses: 0,
	}
	if cs.Winrate != 0 {
		t.Errorf("Winrate should be 0, got %v", cs.Winrate)
	}
}

func TestRecentCards_Initialization(t *testing.T) {
	rc := RecentCards{}
	if rc.BestCard.Winrate != 0 {
		t.Errorf("BestCard winrate should be 0, got %v", rc.BestCard.Winrate)
	}
	if rc.WorstCard.Winrate != 0 {
		t.Errorf("WorstCard winrate should be 0, got %v", rc.WorstCard.Winrate)
	}
	if rc.FrequentCard.Wins != 0 {
		t.Errorf("FrequentCard wins should be 0, got %v", rc.FrequentCard.Wins)
	}
	if rc.FrequentCard.Losses != 0 {
		t.Errorf("FrequentCard losses should be 0, got %v", rc.FrequentCard.Losses)
	}
}
