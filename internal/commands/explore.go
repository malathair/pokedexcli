package commands

import (
	"errors"
	"fmt"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid number of arguments")
	}

	name := args[0]
	location, err := cfg.PokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s . . .\n", location.Name)
	fmt.Println("Pokemon found in this area: ")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf("    %s\n", enc.Pokemon.Name)
	}
	return nil
}
