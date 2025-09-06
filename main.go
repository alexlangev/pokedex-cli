package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type commandRegistry map[string]cliCommand

func getCommands() map[string]cliCommand {
	return commandRegistry{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil // what error should it return
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	args := strings.Split(text, " ")

	return args
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rawInput := ""

	for {
		fmt.Print("Pokedex > ")

		_ = scanner.Scan()
		rawInput = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		cleanedInput := cleanInput(rawInput)
		commands := getCommands()

		switch cleanedInput[0] {
		case commands["exit"].name:
			commands["exit"].callback()

		case commands["help"].name:
			commands["help"].callback()
		}
	}
}
