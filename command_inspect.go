package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if args[0] == "" {
		return fmt.Errorf("Please enter a pokemon!")
	}

	poke := args[0]

	pokemon, ok := cfg.caughtPokemon[poke]
	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf(" -%s: %v\n", s.Stat.Name, s.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" - %s\n", t.Type.Name)
		}
	} else {
		fmt.Printf("you have not caught %s\n", pokemon.Name)
	}

	return nil
}
