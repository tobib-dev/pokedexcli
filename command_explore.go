package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a location name or id")
	}
	areaID := args[0]

	p, err := cfg.pokeapiClient.ListPokemon(areaID)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", p.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range p.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
