package tests

import (
	"net/http"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetLocationAreas(t *testing.T) {
	type fields struct {
		cache  *pokecache.Cache
		client *http.Client
	}
	type args struct {
		url string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    pokeapi.LocationAreaList
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pokeapi.Client{
				Cache:  tt.fields.cache,
				Client: tt.fields.client,
			}
			got, err := c.GetLocationAreas(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocationAreas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocationAreas() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPokemon(t *testing.T) {
	type fields struct {
		cache  *pokecache.Cache
		client *http.Client
	}
	type args struct {
		name string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    pokeapi.Pokemon
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pokeapi.Client{
				Cache:  tt.fields.cache,
				Client: tt.fields.client,
			}
			got, err := c.GetPokemon(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPokemon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPokemonFromArea(t *testing.T) {
	type fields struct {
		cache  *pokecache.Cache
		client *http.Client
	}
	type args struct {
		area string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    pokeapi.PokemonAreaResponse
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pokeapi.Client{
				Cache:  tt.fields.cache,
				Client: tt.fields.client,
			}
			got, err := c.GetPokemonFromArea(tt.args.area)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPokemonFromArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPokemonFromArea() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		cacheInterval time.Duration
	}
	var tests []struct {
		name string
		args args
		want *pokeapi.Client
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pokeapi.NewClient(tt.args.cacheInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
