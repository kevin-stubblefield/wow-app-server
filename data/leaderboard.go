package data

type Leaderboard struct {
	Entries []LeaderboardEntry `json:"entries"`
}

type LeaderboardEntry struct {
	Character             LeaderboardCharacter `json:"character"`
	Rank                  int                  `json:"rank"`
	Rating                int                  `json:"rating"`
	SeasonMatchStatistics MatchStatistics      `json:"season_match_statistics"`
}

type LeaderboardCharacter struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Realm struct {
		Slug string `json:"slug"`
		Id   int    `json:"id"`
	} `json:"realm"`
	Summary *CharacterSummary `json:"summary"`
}

type MatchStatistics struct {
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
}
