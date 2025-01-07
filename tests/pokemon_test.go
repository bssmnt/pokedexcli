package tests

import (
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokemon"
	"reflect"
	"testing"
)

func TestCatchPokemon(t *testing.T) {
	type args struct {
		name   string
		client *pokeapi.Client
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pokemon.CatchPokemon(tt.args.name, tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("CatchPokemon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollectCaughtPokemon(t *testing.T) {
	var tests []struct {
		name string
		want map[string]pokeapi.Pokemon
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pokemon.CollectCaughtPokemon(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectCaughtPokemon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInspect(t *testing.T) {
	type args struct {
		name string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pokemon.Inspect(tt.args.name)
		})
	}
}
