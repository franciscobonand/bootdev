package cli

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/franciscobonand/pokedex/internal/entity"
)

func (c CLI) mapCmd(cfg *entity.Config) Command {
	return Command{
		Name:        "map",
		Description: "Lists 20 locations of the Pokemon world",
		Callback: func() error {
			var val []byte
			url := cfg.NextLocation
			val, next, prev, err := c.handleLocations(url)
			if err != nil {
				return err
			}
			cfg.NextLocation = next
			cfg.PreviousLocation = prev
			c.Cache.Add(url, val)
			return nil
		},
	}
}

func (c CLI) mapbackCmd(cfg *entity.Config) Command {
	return Command{
		Name:        "mapb",
		Description: "Lists previous 20 locations of the Pokemon world",
		Callback: func() error {
			if cfg.PreviousLocation == nil {
				return errors.New("no previous locations")
			}
			url := *cfg.PreviousLocation
			val, next, prev, err := c.handleLocations(url)
			if err != nil {
				return err
			}
			cfg.NextLocation = next
			cfg.PreviousLocation = prev
			c.Cache.Add(url, val)
			return nil
		},
	}
}

func (c CLI) handleLocations(url string) (val []byte, next string, prev *string, err error) {
	if cachedVal, exists := c.Cache.Get(url); !exists {
		resp, err := c.HTTPClient.Get(url)
		if err != nil {
			return nil, "", nil, err
		}
		val = resp
	} else {
		val = cachedVal
	}
	locs := entity.Locations{}
	if err := json.Unmarshal(val, &locs); err != nil {
		return nil, "", nil, err
	}
	for _, l := range locs.Results {
		fmt.Println(l.Name)
	}
	return val, locs.Next, locs.Previous, nil
}
