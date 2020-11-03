package models

type LeaderboardEntry struct {
	Id                 int    `json:"id"`
	Rank               int    `json:"rank"`
	Rating             int    `json:"rating"`
	CharacterName      string `json:"name"`
	CharacterId        int    `json:"character_id"`
	CharacterRealmSlug string `json:"realm_slug"`
	CharacterRealmId   int    `json:"realm_id"`
	CharacterFaction   string `json:"faction"`
	CharacterRace      string `json:"race"`
	CharacterClass     string `json:"class"`
	CharacterSpec      string `json:"spec"`
	GamesPlayed        int    `json:"played"`
	GamesWon           int    `json:"won"`
	GamesLost          int    `json:"lost"`
	Bracket            string `json:"bracket"`
}
