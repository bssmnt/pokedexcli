package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	cache  *pokecache.Cache
	client *http.Client
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache:  pokecache.NewCache(cacheInterval),
		client: &http.Client{},
	}
}

func (c *Client) GetLocationAreas(url string) (LocationAreaList, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	if cached, ok := c.cache.Get(url); ok {
		locationResp := LocationAreaList{}
		err := json.Unmarshal(cached, &locationResp)
		return locationResp, err
	}

	resp, err := c.client.Get(url)
	if err != nil {
		return LocationAreaList{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	c.cache.Add(url, body)

	locationResp := LocationAreaList{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return LocationAreaList{}, err
	}

	return locationResp, nil

}

func (c *Client) GetPokemonFromArea(area string) (PokemonAreaResponse, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL + area

	if cached, ok := c.cache.Get(fullURL); ok {
		locationResp := PokemonAreaResponse{}
		err := json.Unmarshal(cached, &locationResp)
		return locationResp, err
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return PokemonAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	c.cache.Add(fullURL, body)
	locationResp := PokemonAreaResponse{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	return locationResp, err
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	endpoint := "https://pokeapi.co/api/v2/pokemon/" + name

	if data, ok := c.cache.Get(endpoint); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return Pokemon{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(endpoint, data)

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
