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
	callback    func(*config) error
}

type config struct {
	previous string
	next     string
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var commandRegistry map[string]cliCommand

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Displays the names of next 20 location areas in the Pokemon world or first 20 if no previous pages exists",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func startRepl() {
	input := bufio.NewScanner(os.Stdin)

	var cfg config
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		line := cleanInput(input.Text())
		if len(line) == 0 {
			continue
		}
		word := line[0]
		if command, exists := getCommands()[word]; exists == true {
			command.callback(&cfg)
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	newText := strings.TrimSpace(strings.ToLower(text))
	words := strings.Split(newText, " ")

	return words
}
