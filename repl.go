package main

import (
	"strings"
)

func cleanInput(text string) []string {
	newText := strings.TrimSpace(strings.ToLower(text))
	words := strings.Split(newText, " ")

	return words
}
