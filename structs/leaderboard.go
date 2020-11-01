package structs

type Leaderboard struct {
	Entries []struct {
		Character struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		}
	}
}
