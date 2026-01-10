package service

import (
	"clashbot/cr"
	"errors"
)

type CardStats struct {
	Card    cr.Card
	Wins    int
	Losses  int
	Winrate float64
}

type RecentCards struct {
	Cards        []CardStats
	BestCard     CardStats
	WorstCard    CardStats
	FrequentCard CardStats
}

func (s *Service) GetCardStats(tag string) (*RecentCards, error) {
	battleLog, err := s.cr.GetBattleLog(tag)
	if err != nil {
		return nil, err
	}

	cardWins := make(map[cr.Card]int)
	cardLosses := make(map[cr.Card]int)

	for _, battle := range *battleLog {
		team := battle.Team[0]
		opponent := battle.Opponent[0]

		for _, card := range opponent.Cards {
			if team.TrophyChange > 0 {
				cardWins[card]++
			} else {
				cardLosses[card]++
			}
		}
	}

	if len(cardLosses) == 0 && len(cardWins) == 0 {
		return nil, errors.New("недостаточно данных")
	}

	var rc RecentCards

	for card, wins := range cardWins {
		var cs CardStats
		cs = CardStats{
			Card:    card,
			Wins:    wins,
			Losses:  cardLosses[card],
			Winrate: float64(wins) / float64(wins+cardLosses[card]),
		}
		rc.Cards = append(rc.Cards, cs)
	}

	rc.BestCard.Winrate = 0
	rc.WorstCard.Winrate = 1
	rc.FrequentCard.Losses = 0
	rc.FrequentCard.Wins = 0

	for _, card := range rc.Cards {
		if card.Winrate > rc.BestCard.Winrate {
			rc.BestCard = card
		} else if card.Winrate == rc.BestCard.Winrate {
			if (card.Wins + card.Losses) > (rc.BestCard.Wins + rc.BestCard.Losses) {
				rc.BestCard = card
			}
		}

		if card.Winrate < rc.WorstCard.Winrate {
			rc.WorstCard = card
		} else if card.Winrate == rc.BestCard.Winrate {
			if (card.Wins + card.Losses) > (rc.WorstCard.Wins + rc.WorstCard.Losses) {
				rc.WorstCard = card
			}
		}

		if (card.Wins + card.Losses) > (rc.FrequentCard.Wins + rc.FrequentCard.Losses) {
			rc.FrequentCard = card
		}
	}

	return &rc, nil
}
