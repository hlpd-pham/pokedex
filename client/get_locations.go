package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *PokemonClient) GetLocations(url *string) (*NamedAPIResourceList, error) {
	var urlVal string
	var resourceList NamedAPIResourceList

	// sample query: "https://pokeapi.co/api/v2/location/?limit=20&offset=20"
	if url == nil || len(*url) == 0 || *url == "" {
		urlVal = *c.baseUrl + "/location-area"
	} else {
		urlVal = *url
	}

	if cachedData, ok := c.cache.Get(urlVal); ok {
		err := json.Unmarshal(cachedData, &resourceList)
		if err != nil {
			fmt.Printf("Received error while parsing location response: %v\n", err)
			return nil, err
		}
		return &resourceList, nil
	}

	res, err := http.Get(urlVal)
	if err != nil {
		fmt.Printf("Got HTTP Error while requesting location: %v\n", err)
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		return nil, err
	}
	if err != nil {
		fmt.Printf("Received error while requesting location: %v\n", err)
		return nil, err
	}

	err = json.Unmarshal(data, &resourceList)
	if err != nil {
		fmt.Printf("Received error while parsing location response: %v\n", err)
		return nil, err
	}

	c.cache.Add(urlVal, data)
	return &resourceList, nil
}
