package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		var command string
		input.Scan()
		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}
		command = words[0]

		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	newText := strings.TrimSpace(strings.ToLower(text))
	words := strings.Split(newText, " ")

	return words
}
