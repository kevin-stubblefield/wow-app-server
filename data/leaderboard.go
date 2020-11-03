package data

type Leaderboard struct {
	Entries []struct {
		Character struct {
			Name  string `json:"name"`
			Id    int    `json:"id"`
			Realm struct {
				Slug string `json:"slug"`
				Id   int    `json:"id"`
			} `json:"realm"`
			Summary *CharacterSummary `json:"summary"`
		} `json:"character"`
		Rank                  int `json:"rank"`
		Rating                int `json:"rating"`
		SeasonMatchStatistics struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
		} `json:"season_match_statistics"`
	} `json:"entries"`
}
