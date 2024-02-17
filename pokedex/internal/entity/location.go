package entity

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Locations struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type LocationDetails struct {
	ID                int                 `json:"id"`
	Location          Location            `json:"location"`
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
