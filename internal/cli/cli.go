package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/malathair/pokedexcli/internal/commands"
	"github.com/malathair/pokedexcli/internal/config"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func RunCli(cfg *config.Config) {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}

		registry := commands.GetCommandRegistry()
		command, ok := registry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.Callback(cfg, args...); err != nil {
			fmt.Println(err)
		}
	}
}
