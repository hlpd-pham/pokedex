package main

import "fmt"

func commandMap(config *commandConfig) error {
	locationResponse, err := config.PokemonClient.GetLocations(config.NextUrl)
	if err != nil {
		fmt.Printf("Got error while running 'map' command: %v\n", err)
		return nil
	}
	config.NextUrl = &locationResponse.NextUrl
	config.PrevUrl = &locationResponse.PrevUrl

	for _, location := range locationResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapPrev(config *commandConfig) error {
	locationResponse, err := config.PokemonClient.GetLocations(config.PrevUrl)
	if err != nil {
		fmt.Printf("Got error while running 'mapb' command: %v\n", err)
		return nil
	}
	config.NextUrl = &locationResponse.NextUrl
	config.PrevUrl = &locationResponse.PrevUrl

	for _, location := range locationResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
