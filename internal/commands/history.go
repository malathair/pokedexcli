package commands

import (
	"errors"
	"fmt"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandHistory(cfg *config.Config, args ...string) error {
	if len(args) != 0 {
		return errors.New("invalid number of arguments")
	}

	fmt.Println("History:")
	for _, command := range cfg.History.List() {
		fmt.Printf("  %s\n", command)
	}
	
	return nil
}
