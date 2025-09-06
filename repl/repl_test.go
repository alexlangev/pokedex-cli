package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// slice of test case structs
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    "TEST 123 boobs    ",
			expected: []string{"test", "123", "boobs"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len of output is incorrect. Got %d but was expecting %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Incorrect argument, was expecting %s, but got %s", word, expectedWord)
			}
		}
	}
}
