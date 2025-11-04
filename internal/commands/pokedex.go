package commands

import (
	"errors"
	"fmt"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandPokedex(cfg *config.Config, args ...string) error {
	if len(args) != 0 {
		return errors.New("invalid number of arguments")
	}

	if len(cfg.CaughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	fmt.Println("Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
