package cli

import (
	"fmt"

	"github.com/franciscobonand/bootdev/pokedex/internal/entity"
)

func (c CLI) pokedexCmd(cfg *entity.Config) Command {
	return Command{
		Name:        "pokedex",
		Description: "Lists all caught pokemon",
		Callback: func(args ...string) error {
			fmt.Println("Your Pokedex:")
			for name := range cfg.CaughtPokemon {
				fmt.Println(" - ", name)
			}
			return nil
		},
	}
}
