package cli

import "fmt"

func (c CLI) helpCmd() Command {
	return Command{
		Name:        "help",
		Description: "Prints this help message",
		Callback: func() error {
			fmt.Println()
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("Usage:")
			fmt.Println()
			for _, cmd := range c.Commands {
				fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
			}
			fmt.Println()
			return nil
		},
	}
}
