package commands

import (
	"fmt"
	"os"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandExit(cfg *config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}
