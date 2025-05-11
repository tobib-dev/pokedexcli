package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Your pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any pokemon yet, catch a pokemon with the 'catch' command!")
	} else {
		for pokemon, _ := range cfg.caughtPokemon {
			fmt.Println(" - ", pokemon)
		}
	}

	return nil
}
