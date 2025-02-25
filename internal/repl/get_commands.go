package repl

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