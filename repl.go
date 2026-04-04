package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexlangev/pokedex-cli/utils"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays usage message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "return next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "return previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "return the pokemons found at specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "thow a Pokeball to a specified Pokemon",
			callback:    commandCatch,
		},
	}
}

func repl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := utils.CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		arg := ""
		if len(words) >= 2 {
			arg = words[1]
		}
		if command, ok := getCommands()[command]; ok {
			err := command.callback(cfg, arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
