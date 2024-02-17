package cli

import (
	"github.com/franciscobonand/bootdev/pokedex/internal/cache"
	"github.com/franciscobonand/bootdev/pokedex/internal/client"
	"github.com/franciscobonand/bootdev/pokedex/internal/entity"
)

type CLI struct {
	Commands   map[string]Command
	HTTPClient *client.HTTP
	Cache      *cache.Cache
}

type Command struct {
	Name        string
	Description string
	Callback    func(...string) error
	Args        int
}

func NewCLI(httpClient *client.HTTP, c *cache.Cache) *CLI {
	app := &CLI{Commands: make(map[string]Command)}
	cfg := entity.NewConfig()
	app.HTTPClient = httpClient
	app.Cache = c
	app.Commands["help"] = app.helpCmd()
	app.Commands["exit"] = app.exitCmd()
	app.Commands["map"] = app.mapCmd(cfg)
	app.Commands["mapb"] = app.mapbackCmd(cfg)
	app.Commands["explore"] = app.exploreCmd(cfg)
	app.Commands["catch"] = app.catchCmd(cfg)
	app.Commands["inspect"] = app.inspectCmd(cfg)
	return app
}
