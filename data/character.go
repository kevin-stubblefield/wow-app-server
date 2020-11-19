package data

// CharacterSummary holds character specific data
type CharacterSummary struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Realm struct {
		Name string `json:"name"`
	} `json:"realm"`
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

// SpecializationIndex holds spec lsit data
type SpecializationIndex struct {
	CharacterSpecs []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"character_specializations"`
}

// Specialization holds spec specific data
type Specialization struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class struct {
		Name string `json:"name"`
	} `json:"playable_class"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"media"`
}

// SpecIcon holds spec icon data
type SpecIcon struct {
	Assets []struct {
		Value string `json:"value"`
	} `json:"assets"`
}

// ClassIndex holds class list data
type ClassIndex struct {
	Classes []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"classes"`
}

// ClassIcon holds class icon data
type ClassIcon struct {
	Assets []struct {
		Value string `json:"value"`
	} `json:"assets"`
}

// CharacterMedia holds character media urls
type CharacterMedia struct {
	Assets []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"assets"`
}

// Equipment holds equipped items for specific character
type Equipment struct {
	CharacterRealmSlug string
	CharacterName      string
	EquippedItems      []struct {
		Item struct {
			ID int `json:"id"`
		} `json:"item"`
		Slot struct {
			Name string `json:"name"`
		} `json:"slot"`
		Bonuses []int  `json:"bonus_list"`
		Name    string `json:"name"`
	} `json:"equipped_items"`
}
