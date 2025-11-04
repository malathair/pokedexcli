package commands

import (
	"errors"
	"fmt"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandMap(cfg *config.Config, args ...string) error {
	locationsResponse, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResponse.Next
	cfg.PrevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config.Config, args ...string) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	locationsResponse, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResponse.Next
	cfg.PrevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
