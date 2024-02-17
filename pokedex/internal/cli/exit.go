package cli

import (
	"fmt"
	"os"
)

func (c CLI) exitCmd() Command {
	return Command{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback: func(_ ...string) error {
			fmt.Println("Turning off the Pokedex...")
			os.Exit(0)
			return nil
		},
	}
}
