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
		return fmt.Errorf("No pokemons found")
	}

	for _, pokemon := range p.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
