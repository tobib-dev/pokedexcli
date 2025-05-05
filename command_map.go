package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	var apiURL string
	if len(cfg.next) == 0 {
		apiURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		apiURL = cfg.next
	}

	res, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("Error retrieving http response")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var areas LocationAreaResponse
	if err = json.Unmarshal(data, &areas); err != nil {
		return fmt.Errorf("Error unmarshalling response to struct")
	}
	cfg.next = areas.Next

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	var apiURL string

	if len(cfg.previous) == 0 {
		apiURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		apiURL = cfg.previous
	}

	res, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var areas LocationAreaResponse
	if err = json.Unmarshal(data, &areas); err != nil {
		return err
	}

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}
