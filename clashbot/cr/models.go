package cr

type Player struct {
	Tag                             string                   `json:"tag"`
	Name                            string                   `json:"name"`
	ExpLevel                        int                      `json:"expLevel"`
	Trophies                        int                      `json:"trophies"`
	BestTrophies                    int                      `json:"bestTrophies"`
	Wins                            int                      `json:"wins"`
	Losses                          int                      `json:"losses"`
	BattleCount                     int                      `json:"battleCount"`
	ThreeCrownWins                  int                      `json:"threeCrownWins"`
	ChallengeCardsWon               int                      `json:"challengeCardsWon"`
	ChallengeMaxWins                int                      `json:"challengeMaxWins"`
	TournamentCardsWon              int                      `json:"tournamentCardsWon"`
	TournamentBattleCount           int                      `json:"tournamentBattleCount"`
	Role                            string                   `json:"role"`
	Donations                       int                      `json:"donations"`
	DonationsReceived               int                      `json:"donationsReceived"`
	Achievements                    []Achievement            `json:"achievements"`
	TotalDonations                  int                      `json:"totalDonations"`
	WarDayWins                      int                      `json:"warDayWins"`
	ClanCardsCollected              int                      `json:"clanCardsCollected"`
	Clan                            *Clan                    `json:"clan"`
	Arena                           Arena                    `json:"arena"`
	LeagueStatistics                PlayerLeagueStatistics   `json:"leagueStatistics"`
	Badges                          []Badge                  `json:"badges"`
	Cards                           []Card                   `json:"cards"`
	SupportCards                    []Card                   `json:"supportCards"`
	CurrentDeck                     []Card                   `json:"currentDeck"`
	CurrentDeckSupportCards         []Card                   `json:"currentDeckSupportCards"`
	CurrentFavouriteCard            *Card                    `json:"currentFavouriteCard"`
	StarPoints                      int                      `json:"starPoints"`
	ExpPoints                       int                      `json:"expPoints"`
	LegacyTrophyRoadHighScore       int                      `json:"legacyTrophyRoadHighScore"`
	CurrentPathOfLegendSeasonResult PathOfLegendSeasonResult `json:"currentPathOfLegendSeasonResult"`
	LastPathOfLegendSeasonResult    PathOfLegendSeasonResult `json:"lastPathOfLegendSeasonResult"`
	BestPathOfLegendSeasonResult    PathOfLegendSeasonResult `json:"bestPathOfLegendSeasonResult"`
	Progress                        map[string]ProgressEntry `json:"progress"`
	TotalExpPoints                  int                      `json:"totalExpPoints"`
}

type Clan struct {
	Tag              string   `json:"tag"`
	ClanScore        int      `json:"clanScore,omitempty"`
	ClanWarTrophies  int      `json:"clanWarTrophies,omitempty"`
	RequiredTrophies *int     `json:"requiredTrophies,omitempty"`
	DonationsPerWeek int      `json:"donationsPerWeek,omitempty"`
	BadgeId          int      `json:"badgeId"`
	Name             string   `json:"name"`
	Location         Location `json:"location,omitempty"`
	Members          int      `json:"members,omitempty"`
	Description      string   `json:"description,omitempty"`
}

type Arena struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PlayerLeagueStatistics struct {
	PreviousSeason *LeagueSeasonResult `json:"previousSeason,omitempty"`
	BestSeason     *LeagueSeasonResult `json:"bestSeason,omitempty"`
	CurrentSeason  *LeagueSeasonResult `json:"currentSeason,omitempty"`
}

type LeagueSeasonResult struct {
	Trophies     int    `json:"trophies"`
	BestTrophies int    `json:"bestTrophies"`
	Rank         *int   `json:"rank"`
	ID           string `json:"id"`
}

type Badge struct {
	Name     string   `json:"name"`
	Level    int      `json:"level"`
	MaxLevel int      `json:"maxLevel"`
	Progress int      `json:"progress"`
	Target   *int     `json:"target,omitempty"`
	IconUrls IconUrls `json:"iconUrls,omitempty"`
}

type Achievement struct {
	Name           string  `json:"name"`
	Stars          int     `json:"stars"`
	Value          int     `json:"value"`
	Target         *int    `json:"target"`
	Info           string  `json:"info"`
	CompletionInfo *string `json:"completionInfo"`
}

type Card struct {
	Name              string   `json:"name"`
	ID                int      `json:"id"`
	Level             int      `json:"level"`
	EvolutionLevel    int      `json:"evolutionLevel,omitempty"`
	MaxLevel          int      `json:"maxLevel"`
	MaxEvolutionLevel int      `json:"maxEvolutionLevel,omitempty"`
	StarLevel         int      `json:"starLevel,omitempty"`
	Rarity            string   `json:"rarity"`
	Count             int      `json:"count"`
	ElixirCost        int      `json:"elixirCost,omitempty"`
	IconUrls          IconUrls `json:"iconUrls"`
}

type FavouriteCard struct {
	Name              string   `json:"name"`
	ID                int      `json:"id"`
	MaxLevel          int      `json:"maxLevel"`
	MaxEvolutionLevel int      `json:"maxEvolutionLevel,omitempty"`
	Rarity            string   `json:"rarity"`
	ElixirCost        int      `json:"elixirCost"`
	IconUrls          IconUrls `json:"iconUrls"`
}

type PathOfLegendSeasonResult struct {
	LeagueNumber int  `json:"leagueNumber"`
	Trophies     int  `json:"trophies"`
	Rank         *int `json:"rank"`
}

type ProgressEntry struct {
	Arena        Arena `json:"arena"`
	Trophies     int   `json:"trophies"`
	BestTrophies int   `json:"bestTrophies"`
}

type IconUrls struct {
	Large           string `json:"large,omitempty"`
	Medium          string `json:"medium,omitempty"`
	Small           string `json:"small,omitempty"`
	EvolutionMedium string `json:"evolutionMedium,omitempty"`
}

type Location struct {
	LocalizedName string `json:"localizedName"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	IsCountry     bool   `json:"isCountry"`
	CountryCode   string `json:"countryCode"`
}
