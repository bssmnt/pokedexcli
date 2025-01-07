package pokeapi

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type PokemonAreaResponse struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type StatDetail struct {
	Name string `json:"name"`
}

type Stat struct {
	BaseStat int        `json:"base_stat"`
	Stat     StatDetail `json:"stat"`
}

type TypeDetail struct {
	Name string `json:"name"`
}

type TypeInfo struct {
	Slot int        `json:"slot"`
	Type TypeDetail `json:"type"`
}

type Pokemon struct {
	Name           string     `json:"name"`
	BaseExperience int        `json:"base_experience"`
	Height         int        `json:"height"`
	Weight         int        `json:"weight"`
	Stats          []Stat     `json:"stats"`
	Types          []TypeInfo `json:"types"`
	URL            string     `json:"url"`
}
