package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *PokemonClient) CatchPokemon(pokemonName *string) (*Pokemon, error) {
	var urlVal string
	var pokemon Pokemon

	if pokemonName == nil || len(*pokemonName) == 0 {
		return &pokemon, errors.New("pokemonName is empty")
	}

	urlVal = fmt.Sprintf("%s/pokemon/%s", *c.baseUrl, *pokemonName)

	if cachedData, ok := c.cache.Get(urlVal); ok {
		fmt.Printf("Found cached entry for url: %s\n", urlVal)
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			fmt.Printf("Received error while parsing pokemon response: %v\n", err)
			return nil, err
		}
		return &pokemon, nil
	}

	fmt.Printf("Cache miss, requesting entry for url: %s\n", urlVal)
	res, err := http.Get(urlVal)
	if err != nil {
		fmt.Printf("Got HTTP Error while requesting pokemon: %s, err: %v\n", *pokemonName, err)
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		return nil, err
	}
	if err != nil {
		fmt.Printf("Received error while requesting pokemon: %v\n", err)
		return nil, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Printf("Received error while parsing location response: %v\n", err)
		return nil, err
	}

	c.cache.Add(urlVal, data)
	return &pokemon, nil
}
