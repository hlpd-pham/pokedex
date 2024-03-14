package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hlpd-pham/pokedexcli/client"
)

type commandConfig struct {
	PokemonClient client.PokemonClient
	PrevUrl       *string
	NextUrl       *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *commandConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display 20 locations from next page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display 20 locations from previous page",
			callback:    commandMapPrev,
		},
	}
}

func parseInput(text string) []string {
	lowercase := strings.ToLower(text)
	trimmed := strings.Trim(lowercase, " ")
	tokens := strings.Split(trimmed, " ")
	return tokens
}

func main() {
	cmdCfg := commandConfig{
		PokemonClient: client.NewPokemonClient(),
		PrevUrl:       nil,
		NextUrl:       nil,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputTokens := parseInput(scanner.Text())
		command, ok := getCommands()[inputTokens[0]]
		if !ok {
			fmt.Println("Unknown command")
			commandHelp(&cmdCfg)
		} else {
			command.callback(&cmdCfg)
		}
		fmt.Println()
	}
}
