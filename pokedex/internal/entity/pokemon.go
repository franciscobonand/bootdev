package entity

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}
