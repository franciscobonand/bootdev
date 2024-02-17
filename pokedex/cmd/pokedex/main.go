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
		cmds := sanitizeInput(input.Text())
		if len(cmds) == 0 {
			continue
		}
		cmd := cmds[0]
		args := cmds[1:]
		if c, ok := app.Commands[cmd]; ok {
			if len(args) != c.Args {
				fmt.Printf("Command '%s' requires %d arguments\n", cmd, c.Args)
				continue
			}
			if err := c.Callback(args...); err != nil {
				fmt.Println("Error executing command:", err)
			}
			continue
		}
		fmt.Printf("Unknown command '%s'\n", cmds)
		app.Commands["help"].Callback()
	}
}

func sanitizeInput(s string) []string {
	cmds := []string{}
	for _, cmd := range strings.Fields(s) {
		cmds = append(cmds, strings.ToLower(cmd))
	}
	return cmds
}
