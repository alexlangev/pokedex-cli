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
	callback    func(*Config) error
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
		if command, ok := getCommands()[command]; ok {
			err := command.callback(cfg)
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
