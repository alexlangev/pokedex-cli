package utils

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "My name IS",
			expected: []string{"my", "name", "is"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match, expected %v and got %v", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words don't match, expected %v and got %v", expectedWord, word)
			}
		}
	}
}
