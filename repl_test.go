package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  I'm sorry I hurt you ",
			expected: []string{"i'm", "sorry", "i", "hurt", "you"},
		},
		{
			input:    "Thanks for the other day",
			expected: []string{"thanks", "for", "the", "other", "day"},
		},
		{
			input:    " Let's dance till the end of time",
			expected: []string{"let's", "dance", "till", "the", "end", "of", "time"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test Failed => actual length: %d is not same as expected length: %d", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Test Failed => word: %s is not same as expected word: %s", word, expectedWord)
			}
		}
	}
}
