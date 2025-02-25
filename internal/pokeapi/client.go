package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	Cache  *pokecache.Cache
	Client *http.Client
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		Cache:  pokecache.NewCache(cacheInterval),
		Client: &http.Client{},
	}
}

func (c *Client) GetLocationAreas(url string) (LocationAreaList, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	if cached, ok := c.Cache.Get(url); ok {
		locationResp := LocationAreaList{}
		err := json.Unmarshal(cached, &locationResp)
		return locationResp, err
	}

	resp, err := c.Client.Get(url)
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

	c.Cache.Add(url, body)

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

	if cached, ok := c.Cache.Get(fullURL); ok {
		locationResp := PokemonAreaResponse{}
		err := json.Unmarshal(cached, &locationResp)
		return locationResp, err
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	c.Cache.Add(fullURL, body)
	locationResp := PokemonAreaResponse{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	return locationResp, err
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	endpoint := "https://pokeapi.co/api/v2/pokemon/" + name

	if data, ok := c.Cache.Get(endpoint); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := c.Client.Get(endpoint)
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

	c.Cache.Add(endpoint, data)

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
