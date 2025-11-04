package main

import (
	"time"

	"github.com/malathair/pokedexcli/internal/cli"
	"github.com/malathair/pokedexcli/internal/config"
	"github.com/malathair/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config.Config{
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
		PokeapiClient: pokeClient,
	}

	cli.RunCli(cfg)
}
