package main

type Config struct {
	BotToken    string `json:"BotToken"`
	ClashApiKey string `json:"ClashApiKey"`
}

type Player struct {
	Tag            string `json:"tag"`
	Name           string `json:"name"`
	ExpLevel       int    `json:"expLevel"`
	Trophies       int    `json:"trophies"`
	BestTrophies   int    `json:"bestTrophies"`
	Wins           int    `json:"wins"`
	Losses         int    `json:"losses"`
	BattleCount    int    `json:"battleCount"`
	ThreeCrownWins int    `json:"threeCrownWins"`
}
