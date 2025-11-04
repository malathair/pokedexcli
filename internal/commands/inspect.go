package commands

import (
	"errors"
	"fmt"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	name := args[0]

	pokemon, ok := cfg.CaughtPokemon[name]
	if !ok {
		return errors.New("Pokemon not caught")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
