package cli

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"

	"github.com/franciscobonand/bootdev/pokedex/internal/entity"
)

const (
	catchPath = "/pokemon/%s"
	// catchChanceDelimiter is an arbitrary number used to calculate the chance of catching a pokemon
	catchChanceDelimiter = 45
)

func (c CLI) catchCmd(cfg *entity.Config) Command {
	return Command{
		Name:        "catch",
		Description: "Tries to catch a pokemon",
		Args:        1,
		Callback: func(args ...string) error {
			url := fmt.Sprintf(cfg.BaseURL+catchPath, args[0])
			val, pkmon, err := c.handlePokemonCatch(url)
			if err != nil {
				return err
			}
			if pkmon != nil {
				cfg.CaughtPokemon[pkmon.Name] = *pkmon
			}
			c.Cache.Add(url, val)
			return nil
		},
	}
}

func (c CLI) handlePokemonCatch(url string) ([]byte, *entity.PokemonDetails, error) {
	var val []byte
	if cachedVal, exists := c.Cache.Get(url); !exists {
		resp, err := c.HTTPClient.Get(url)
		if err != nil {
			return nil, nil, err
		}
		val = resp
	} else {
		val = cachedVal
	}
	pkmonDetails := entity.PokemonDetails{}
	if err := json.Unmarshal(val, &pkmonDetails); err != nil {
		return nil, nil, err
	}
	fmt.Printf("Throwing a pokeball at %s...\n", pkmonDetails.Name)
	chance := rand.IntN(pkmonDetails.BaseExperience)
	if chance < catchChanceDelimiter {
		fmt.Printf("%s was caught!\n", pkmonDetails.Name)
		return val, &pkmonDetails, nil
	}
	fmt.Printf("%s ran away!\n", pkmonDetails.Name)
	return val, nil, nil
}
