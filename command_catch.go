package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Enter the pokemon you want to catch!")
	}

	pokeName := args[0]
	pokemon, err := cfg.pokeapiClient.PokemonGet(pokeName)
	if err != nil {
		return err
	}

	catchChance := 100 - (pokemon.BaseExperience / 4)
	if catchChance < 5 {
		catchChance = 5
	}
	if catchChance > 90 {
		catchChance = 90
	}

	catch := rand.Intn(100) + 1
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catch <= catchChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
