package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	areas, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = areas.Next
	cfg.prevLocationsURL = areas.Previous

	for _, loc := range areas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	areas, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = areas.Next
	cfg.prevLocationsURL = areas.Previous

	for _, loc := range areas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
