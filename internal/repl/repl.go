package repl

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokemon"
	"strings"
	"time"
)

type CliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args ...string) error
}

type Config struct {
	Next          string
	Previous      string
	pokeapiClient *pokeapi.Client
}

func REPL() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		pokeapiClient: pokeapi.NewClient(5 * time.Minute),
	}

	err := CommandHelp(cfg)
	if err != nil {
		return
	}

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		cleanInput := CleanInput(scanner.Text())
		if len(cleanInput) == 0 {
			continue
		}

		commandName := cleanInput[0]
		args := cleanInput[1:]

		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}

	}
}

func CleanInput(text string) []string {
	cleanText := strings.TrimSpace(text)
	lowerCase := strings.ToLower(cleanText)
	splitText := strings.Fields(lowerCase)
	return splitText
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 areas",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 areas",
			callback:    CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    CallbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    CommandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your caught Pokemon",
			callback:    CommandPokedex,
		},
		"inspect": {
			name:        "inspect",
			description: "Display pokemon information",
			callback:    CommandInspect,
		},
	}
}

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
