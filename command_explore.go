package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *commandConfig) error {
	if len(config.CommandArgs) == 0 {
		return errors.New("need location name for explore command")
	}
	if len(config.CommandArgs) > 1 {
		return errors.New("only accept 1 location for explore command")
	}
	fmt.Printf("Exploring %s...\n", config.CommandArgs[0])
	locationResponse, err := config.PokemonClient.ExploreLocation(&config.CommandArgs[0])
	if err != nil || locationResponse == nil {
		fmt.Printf("Got error during 'explore' command: %v, or location response is nil\n", err)
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationResponse.PokemonEncounters {
		fmt.Printf("  - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
