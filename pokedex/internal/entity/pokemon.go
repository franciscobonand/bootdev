package entity

import "fmt"

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type PokemonDetails struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

func (p PokemonDetails) String() string {
	var stats, types string
	for _, s := range p.Stats {
		stats += fmt.Sprintf(" - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	for _, t := range p.Types {
		types += fmt.Sprintf(" - %s\n", t.Type.Name)
	}
	return "\nName: " + p.Name + "\n" +
		"Height: " + fmt.Sprint(p.Height) + "\n" +
		"Weight: " + fmt.Sprint(p.Weight) + "\n" +
		"Stats: \n" + stats +
		"Types: \n" + types
}
