package config

import (
	"github.com/malathair/pokedexcli/internal/history"
	"github.com/malathair/pokedexcli/internal/pokeapi"
)

type Config struct {
	CaughtPokemon    map[string]pokeapi.Pokemon
	History          history.History
	NextLocationsURL *string
	PrevLocationsURL *string
	PokeapiClient    *pokeapi.Client
}
