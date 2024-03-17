package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *commandConfig) error {
	if len(config.CommandArgs) == 0 {
		return errors.New("need location name for explore command")
	}
	if len(config.CommandArgs) > 1 {
		return errors.New("only accept 1 pokemon for inspect command")
	}

	if pokemon, ok := config.PokemonCaught[config.CommandArgs[0]]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}
	} else {
		fmt.Printf("Pokemon %s has not been caught\n", config.CommandArgs[0])
	}
	return nil
}
