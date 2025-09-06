package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	args := strings.Split(text, " ")

	return args
}

func main() {
	for {
		fmt.Print("Pokedex > ")

		scanner := bufio.NewScanner(os.Stdin)
		rawInput := ""

		_ = scanner.Scan()
		rawInput = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		cleanedInput := cleanInput(rawInput)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}
