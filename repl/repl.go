package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/alexlangev/pokedex-cli/client"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type commandRegistry map[string]cliCommand

type config struct {
	nextUrl string
	prevUrl string
	client  client.Client
}

type ResponseLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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
		"map": {
			name:        "map",
			description: "Print the next 20 locations",
			callback:    commandMap,
		},
	}
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil // what error should it return
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	args := strings.Split(text, " ")

	return args
}

func commandMap(c *config) error {
	var url string

	if c.nextUrl == "" && c.prevUrl == "" {
		url = c.client.BaseURL + "location-area/"
	} else if c.nextUrl == "" {
		fmt.Println("You've reached the last page")
		return nil
	} else {
		url = c.nextUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := c.client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locations ResponseLocations
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	if locations.Next != nil {
		c.nextUrl = *locations.Next
	} else {
		c.nextUrl = ""
	}
	if locations.Previous != nil {
		c.prevUrl = *locations.Previous
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func Repl() {
	cli := client.NewClient()
	cfg := config{
		client: *cli,
	}

	scanner := bufio.NewScanner(os.Stdin)
	rawInput := ""

	for {
		fmt.Print("Pokedex > ")

		_ = scanner.Scan()
		rawInput = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		commands := getCommands()
		cleanedInput := cleanInput(rawInput)
		if len(cleanedInput) < 1 {
			continue
		}

		switch cleanedInput[0] {
		case commands["exit"].name:
			commands["exit"].callback(&cfg)

		case commands["help"].name:
			commands["help"].callback(&cfg)

		case commands["map"].name:
			commands["map"].callback(&cfg)
		}
	}
}
