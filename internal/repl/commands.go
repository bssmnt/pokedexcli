package repl

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokemon")

func CommandHelp(*Config, ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func CommandExit(*Config, ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandMap(cfg *Config, _ ...string) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func CommandMapB(cfg *Config, _ ...string) error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page!")
		return nil
	}

	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func CallbackExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <area-name>")
	}

	poke, err := cfg.pokeapiClient.GetPokemonFromArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, encounter := range poke.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func CommandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <name>")
	}

	return pokemon.CatchPokemon(args[0], cfg.pokeapiClient)
}

func CommandPokedex(*Config, ...string) error {
	caught := pokemon.CollectCaughtPokemon()
	if len(caught) == 0 {
		fmt.Println("Your Pokedex is still empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, p := range caught {
		fmt.Printf("- %s\n", p.Name)
	}

	return nil
}

func CommandInspect(_ *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide the name of a Pokemon to inspect")
	}

	name := args[0]

	pokemon.Inspect(name)

	return nil
}