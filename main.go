package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	args := strings.Split(text, " ")

	return args
}

func main() {
	fmt.Println("Hello, World!")
}
