package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(*config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
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

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}
	return nil
}
