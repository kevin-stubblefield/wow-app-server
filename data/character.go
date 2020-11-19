package data

type CharacterSummary struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Faction struct {
		Name string `json:"name"`
	} `json:"faction"`
	Race struct {
		Name string `json:"name"`
	} `json:"race"`
	Class struct {
		Name string `json:"name"`
	} `json:"character_class"`
	Spec struct {
		Name string `json:"name"`
	} `json:"active_spec"`
}

type Equipment struct {
	EquippedItems []struct {
		Item struct {
			Id int `json:"id"`
		} `json:"item"`
		Slot struct {
			Name string `json:"name"`
		} `json:"slot"`
		Bonuses []int  `json:"bonus_list"`
		Name    string `json:"name"`
	} `json:"equipped_items"`
}
