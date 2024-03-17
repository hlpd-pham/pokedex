package main

import "fmt"

func commandPokedex(config *commandConfig) error {
	if len(config.PokemonCaught) == 0 {
		fmt.Printf("Your Pokedex is empty\n")
		return nil
	}
	fmt.Printf("Your Pokedex:\n")
	for pokemonName := range config.PokemonCaught {
		fmt.Printf("  - %s\n", pokemonName)
	}
	return nil
}
