package entity

type Config struct {
	BaseURL          string
	NextLocation     string
	PreviousLocation *string
	CaughtPokemon    map[string]PokemonDetails
}

func NewConfig() *Config {
	return &Config{
		BaseURL:          "https://pokeapi.co/api/v1",
		NextLocation:     "https://pokeapi.co/api/v1/location-area?offset=0&limit=20",
		PreviousLocation: nil,
		CaughtPokemon:    map[string]PokemonDetails{},
	}
}
