package data

// Leaderboard holds list of leaderboard entries
type Leaderboard struct {
	Entries []LeaderboardEntry `json:"entries"`
}

// LeaderboardEntry holds data for an entry in a leaderboard
type LeaderboardEntry struct {
	Character             LeaderboardCharacter `json:"character"`
	Rank                  int                  `json:"rank"`
	Rating                int                  `json:"rating"`
	SeasonMatchStatistics MatchStatistics      `json:"season_match_statistics"`
}

// LeaderboardCharacter holds character data within the leaderboard
type LeaderboardCharacter struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Realm struct {
		Slug string `json:"slug"`
		ID   int    `json:"id"`
	} `json:"realm"`
	Summary CharacterSummary `json:"summary"`
}

// MatchStatistics holds win loss played values
type MatchStatistics struct {
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
}
