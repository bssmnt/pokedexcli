package repl

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
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