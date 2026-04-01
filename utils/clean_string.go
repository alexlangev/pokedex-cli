package utils

import (
	"strings"
)

// Split user input in "words" based on whitespace
// no need to handle punctuation, this is for cmdline arguments
func CleanInput(text string) []string {
	cleanString := strings.ToLower(text)
	cleanWords := strings.Fields(cleanString)

	return cleanWords
}
