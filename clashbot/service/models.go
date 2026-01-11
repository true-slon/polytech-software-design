package service

import (
	"clashbot/cr"
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
