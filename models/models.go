package models

// LeaderboardEntry holds data for each entry in the leaderboard table
type LeaderboardEntry struct {
	ID                 int    `json:"id"`
	Rank               int    `json:"rank"`
	Rating             int    `json:"rating"`
	CharacterName      string `json:"name"`
	CharacterID        int    `json:"character_id"`
	CharacterRealm     string `json:"realm"`
	CharacterRealmSlug string `json:"realm_slug"`
	CharacterRealmID   int    `json:"realm_id"`
	CharacterFaction   string `json:"faction"`
	CharacterRace      string `json:"race"`
	CharacterClass     string `json:"class"`
	CharacterSpec      string `json:"spec"`
	GamesPlayed        int    `json:"played"`
	GamesWon           int    `json:"won"`
	GamesLost          int    `json:"lost"`
	Bracket            string `json:"bracket"`
}

// Character holds data for each entry in the character table
type Character struct {
	Name      string `json:"name"`
	Realm     string `json:"realm"`
	RealmSlug string `json:"realm_slug"`
}

// Specialization holds data for each entry in the specialization table
type Specialization struct {
	ID        int    `json:"id"`
	Class     string `json:"class"`
	ClassSlug string `json:"class_slug"`
	Spec      string `json:"spec"`
	SpecSlug  string `json:"spec_slug"`
	SpecRole  string `json:"spec_role"`
}
