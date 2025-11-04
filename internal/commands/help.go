package commands

import (
	"fmt"
	"sort"

	"github.com/malathair/pokedexcli/internal/config"
)

func commandHelp(cfg *config.Config) error {
	fmt.Println()
	fmt.Println("Usage:")

	registry := GetCommandRegistry()
	commands := make([]string, 0, len(registry))
	for k := range registry {
		commands = append(commands, k)
	}
	sort.Strings(commands)

	for _, command := range commands {
		fmt.Printf("  %-10s %s\n", registry[command].Name, registry[command].Description)
	}

	fmt.Println()
	return nil
}
