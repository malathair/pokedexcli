package commands

import (
	"github.com/malathair/pokedexcli/internal/config"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config, ...string) error
}

func GetCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			Name:        "catch <pokemon>",
			Description: "Catch a pokemon",
			Callback:    commandCatch,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"explore": {
			Name:        "explore <area>",
			Description: "Explore an area",
			Callback:    commandExplore,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"inspect": {
			Name:        "inspect <pokemon>",
			Description: "View details about a captured pokemon",
			Callback:    commandInspect,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapb,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List all pokemon in the Pokedex",
			Callback:    commandPokedex,
		},
	}
}
