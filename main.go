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
	CommandArgs   []string
	PokemonCaught map[string]client.Pokemon
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
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Reveal your pokedex",
			callback:    commandPokedex,
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
		PokemonCaught: make(map[string]client.Pokemon),
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
			cmdCfg.CommandArgs = inputTokens[1:]
			err := command.callback(&cmdCfg)
			if err != nil {
				fmt.Printf("Found error while running command %s: %v\n", command.name, err)
			}
		}
		fmt.Println()
	}
}
