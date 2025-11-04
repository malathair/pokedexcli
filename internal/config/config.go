package config

import "github.com/malathair/pokedexcli/internal/pokeapi"

type Config struct {
	PokeapiClient    *pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}
