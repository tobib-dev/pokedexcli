package main

import (
	"time"

	"github.com/tobib-dev/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.PokemonTraitsResponse{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
