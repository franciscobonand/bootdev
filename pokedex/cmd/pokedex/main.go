package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/franciscobonand/bootdev/pokedex/internal/cache"
	cli "github.com/franciscobonand/bootdev/pokedex/internal/cli"
	"github.com/franciscobonand/bootdev/pokedex/internal/client"
)

func main() {
	httpClient := client.NewHTTPClient()
	c := cache.New(5 * time.Second)
	app := cli.NewCLI(httpClient, c)
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		input.Scan()
		if err := input.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		cmd := sanitizeInput(input.Text())
		if c, ok := app.Commands[cmd]; ok {
			if err := c.Callback(); err != nil {
				fmt.Println("Error executing command:", err)
			}
			continue
		}
		fmt.Printf("Unknown command '%s'\n", cmd)
		app.Commands["help"].Callback()
	}
}

func sanitizeInput(s string) string {
	cmd := strings.Fields(s)[0]
	cmd = strings.ToLower(cmd)
	return cmd
}
