package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandRegistry map[string]cliCommand

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, com := range commandRegistry {
		name := com.name
		des := com.description
		fmt.Printf("%s: %s\n", name, des)
	}

	return nil
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func init() {
	commandRegistry = map[string]cliCommand{
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
	}
}

func startRepl() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		line := cleanInput(input.Text())
		if len(line) == 0 {
			continue
		}
		word := line[0]
		if command, exists := commandRegistry[word]; exists == true {
			command.callback()
		}
		fmt.Print("Unknown command\n")
	}
}

func cleanInput(text string) []string {
	newText := strings.TrimSpace(strings.ToLower(text))
	words := strings.Split(newText, " ")

	return words
}
