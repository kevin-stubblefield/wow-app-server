package structs

type SeasonIndex struct {
	CurrentSeason struct {
		Id int `json:"id"`
	} `json:"current_season"`
}
