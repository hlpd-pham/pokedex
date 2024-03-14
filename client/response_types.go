package client

// https://pokeapi.co/docs/v2#namedapiresource
type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// https://pokeapi.co/docs/v2#location-areas
type NamedAPIResourceList struct {
	Count   int                `json:"count"`
	NextUrl string             `json:"next"`
	PrevUrl string             `json:"previous"`
	Results []NamedAPIResource `json:"results"`
}
