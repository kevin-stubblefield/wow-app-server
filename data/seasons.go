package data

// SeasonIndex holds data for the current season
type SeasonIndex struct {
	CurrentSeason struct {
		ID int `json:"id"`
	} `json:"current_season"`
}
