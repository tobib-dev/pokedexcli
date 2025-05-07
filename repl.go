package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tobib-dev/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	prevLocationsURL *string
	nextLocationsURL *string
}

var commandRegistry map[string]cliCommand

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world or first 20 if no previous pages exists",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokemon located in an area",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(cfg *config) {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		line := cleanInput(input.Text())

		if len(line) == 0 {
			continue
		}
		word := line[0]
		option := ""
		if len(line) > 1 {
			option = strings.Join(line[1:], "-")
		}

		if command, exists := getCommands()[word]; exists == true {
			err := command.callback(cfg, option)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	newText := strings.TrimSpace(strings.ToLower(text))
	words := strings.Split(newText, " ")

	return words
}
