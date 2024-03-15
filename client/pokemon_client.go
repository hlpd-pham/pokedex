package client

import (
	"time"

	"github.com/hlpd-pham/pokedexcli/cache"
)

type Client interface {
	GetLocations(limit, offset int) error
}

type PokemonClient struct {
	Client
	baseUrl *string
	cache   *cache.Cache
}

func NewPokemonClient() PokemonClient {
	rootUrl := "https://pokeapi.co/api/v2"
	return PokemonClient{
		baseUrl: &rootUrl,
		cache:   cache.NewCache(30 * time.Second),
	}
}
