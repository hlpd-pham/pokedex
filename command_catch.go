package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *commandConfig) error {
	if len(config.CommandArgs) == 0 {
		return errors.New("need Pokemon name for catch command")
	}
	if len(config.CommandArgs) > 1 {
		return errors.New("only accept 1 Pokemon for catch command")
	}
	fmt.Printf("Catching %s...\n", config.CommandArgs[0])
	pokemonResponse, err := config.PokemonClient.CatchPokemon(&config.CommandArgs[0])
	if err != nil || pokemonResponse == nil {
		fmt.Printf("Got error during 'catch' command: %v, or pokemon response is nil\n", err)
		return nil
	}

	playerRoll := rand.Intn(pokemonResponse.BaseExperience)

	fmt.Printf("Found Pokemon: %s, baseEXP: %d, player rolled: %d\n",
		pokemonResponse.Name, pokemonResponse.BaseExperience, playerRoll)

	if playerRoll < pokemonResponse.BaseExperience/2 {
		fmt.Printf("%s ran away!", pokemonResponse.Name)
		return nil
	}

	fmt.Printf("%s caught!", pokemonResponse.Name)
	config.PokemonCaught[pokemonResponse.Name] = *pokemonResponse
	return nil
}
