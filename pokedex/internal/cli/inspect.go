package cli

import (
	"fmt"

	"github.com/franciscobonand/bootdev/pokedex/internal/entity"
)

func (c CLI) inspectCmd(cfg *entity.Config) Command {
	return Command{
		Name:        "inspect",
		Description: "Shows details of a pokemon that has been caught",
		Args:        1,
		Callback: func(args ...string) error {
			name := args[0]
			if pkmon, exists := cfg.CaughtPokemon[name]; exists {
				fmt.Println(pkmon)
				return nil
			}
			fmt.Println("You have not caught that pokemon yet!")
			return nil
		},
	}
}
