package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *PokemonClient) ExploreLocation(locationName *string) (*LocationArea, error) {
	var urlVal string
	var locationArea LocationArea

	if locationName == nil || len(*locationName) == 0 {
		return &locationArea, errors.New("locationName is empty")
	}

	urlVal = fmt.Sprintf("%s/location-area/%s", *c.baseUrl, *locationName)

	if cachedData, ok := c.cache.Get(urlVal); ok {
		fmt.Printf("Found cached entry for url: %s\n", urlVal)
		err := json.Unmarshal(cachedData, &locationArea)
		if err != nil {
			fmt.Printf("Received error while parsing location-area response: %v\n", err)
			return nil, err
		}
		return &locationArea, nil
	}

	fmt.Printf("Cache miss, requesting entry for url: %s\n", urlVal)
	res, err := http.Get(urlVal)
	if err != nil {
		fmt.Printf("Got HTTP Error while requesting location-area: %s, err: %v\n", *locationName, err)
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

	// fmt.Printf("Found data for location: %s - %s\n", *locationName, data)
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		fmt.Printf("Received error while parsing location response: %v\n", err)
		return nil, err
	}

	c.cache.Add(urlVal, data)
	return &locationArea, nil
}
