package cli

import (
	"encoding/json"
	"fmt"

	"github.com/franciscobonand/bootdev/pokedex/internal/entity"
)

const (
	explorePath = "/location-area/%s"
)

func (c CLI) exploreCmd(cfg *entity.Config) Command {
	return Command{
		Name:        "explore",
		Description: "Lists pokemon found in a given area",
		Args:        1,
		Callback: func(args ...string) error {
			url := fmt.Sprintf(cfg.BaseURL+explorePath, args[0])
			val, err := c.handlePokemonList(url)
			if err != nil {
				return err
			}
			c.Cache.Add(url, val)
			return nil
		},
	}
}

func (c CLI) handlePokemonList(url string) ([]byte, error) {
	var val []byte
	if cachedVal, exists := c.Cache.Get(url); !exists {
		resp, err := c.HTTPClient.Get(url)
		if err != nil {
			return nil, err
		}
		val = resp
	} else {
		val = cachedVal
	}
	locDetails := entity.LocationDetails{}
	if err := json.Unmarshal(val, &locDetails); err != nil {
		return nil, err
	}
	fmt.Println("Pokemon found in this area:")
	for _, e := range locDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", e.Pokemon.Name)
	}
	return val, nil
}
