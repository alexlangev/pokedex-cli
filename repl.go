package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func repl() {
	const prompt = "Pokedex > "

	scanner := bufio.NewScanner(os.Stdin)

	config := config{
		prevURL: "",
		nextURL: "https://pokeapi.co/api/v2/location-area/?offset=0",
	}

	for {
		fmt.Print(prompt)
		// block and wait for user input
		scanner.Scan()

		userInput := scanner.Text()
		args := cleanInput(userInput)

		if len(args) == 0 {
			continue
		}

		// does command exist? is so call the callback
		if cmd, ok := commands[args[0]]; ok {
			cmd.callback(&config)
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Println("        ")
		fmt.Println("next: ", config.nextURL)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextURL string
	prevURL string
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
	"map": {
		name:        "map",
		description: "display the next 20 areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "display the previous 20 areas",
		callback:    commandMapb,
	},
}

func commandExit(*config) error {
	const msg = "Closing the Pokedex... Goodbye!"
	fmt.Println(msg)
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	const msg = "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex"
	fmt.Println(msg)
	return nil
}

func commandMap(c *config) error {
	if c.nextURL == "" {
		fmt.Println("Already at the end")
		return nil
	}

	res, err := http.Get(c.nextURL)
	if err != nil {
		return err
	}

	// read response
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// unmarshal
	var data locationAreas
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	// update config with new next and prev
	c.prevURL = data.Previous
	c.nextURL = data.Next

	// print location areas
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.prevURL == "" {
		fmt.Println("Already at the start")
		return nil
	}
	res, err := http.Get(c.prevURL)
	if err != nil {
		return err
	}

	// read response
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// unmarshal
	var data locationAreas
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	// update config with new next and prev
	c.prevURL = data.Previous
	c.nextURL = data.Next

	// print location areas
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

type locationAreas struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
