package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	name := args[0]
	pokemon, err := cfg.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	result := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if result > 60 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.CaughtPokemon[pokemon.Name] = pokemon

	return nil
}
