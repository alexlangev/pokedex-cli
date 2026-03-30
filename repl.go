package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Split user input in "words" based on whitespace
// no need to handle punctuation, this is for cmdline arguments
func cleanInput(text string) []string {
	cleanString := strings.ToLower(text)
	cleanWords := strings.Fields(cleanString)

	return cleanWords
}

// infinite loop until exit
// prints on screen
// reads stdin
// repeat
func repl() {
	const prompt = "Pokedex > "

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		// block and wait for user input
		scanner.Scan()

		userInput := scanner.Text()
		args := cleanInput(userInput)

		// does command exist? is so call the callback
		if cmd, ok := commands[args[0]]; ok {
			cmd.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "displays a help message",
		callback:    commandHelp,
	},
}

func commandExit() error {
	const msg = "Closing the Pokedex... Goodbye!"
	fmt.Println(msg)
	os.Exit(0)
	return nil
}

func commandHelp() error {
	const msg = "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex"
	fmt.Println(msg)
	return nil
}
