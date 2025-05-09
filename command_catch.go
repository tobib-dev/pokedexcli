package main

import (
	"fmt"
	"math/rand"

	"github.com/tobib-dev/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Enter the pokemon you want to catch!")
	}

	pokeName := args[0]
	p, err := cfg.pokeapiClient.PokemonGet(pokeName)
	var pokedex map[string]pokeapi.PokemonTraitsResponse
	if err != nil {
		return err
	}

	catchChance := 100 - (p.BaseExperience / 4)
	if catchChance < 5 {
		catchChance = 5
	}
	if catchChance > 90 {
		catchChance = 90
	}

	catch := rand.Intn(100) + 1
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	if catch <= catchChance {
		fmt.Printf("%s was caught!\n", pokeName)
		if len(pokedex) == 0 {
			pokedex := make(map[string]pokeapi.PokemonTraitsResponse)
			pokedex[pokeName] = p
		} else {
			pokedex[pokeName] = p
		}
	} else {
		fmt.Printf("%s escaped!\n", pokeName)
	}

	return nil
}
