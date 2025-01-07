package pokemon

import (
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokeapi"
)

var pokedex map[string]pokeapi.Pokemon

func init() {
	pokedex = make(map[string]pokeapi.Pokemon)
}

func CatchPokemon(name string, client *pokeapi.Client) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := client.GetPokemon(name)
	if err != nil {
		return err
	}

	catchRoll := rand.Intn(pokemon.BaseExperience * 3)

	if catchRoll < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", name)
	} else {
		pokedex[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
	}

	return nil
}

func CollectCaughtPokemon() map[string]pokeapi.Pokemon {
	return pokedex
}

func Inspect(name string) {
	pokemon, found := pokedex[name]
	if found {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf("- %s\n", typ.Type.Name)
		}
	} else {
		fmt.Printf("You haven't caught %s!\n", name)
	}
}
