package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(areaID string) (PokemonAreaResponse, error) {
	url := baseURL + "/location-area/" + areaID

	cacheData, exists := c.cache.Get(url)
	if exists {
		pokemonList := PokemonAreaResponse{}
		err := json.Unmarshal(cacheData, &pokemonList)
		if err != nil {
			return PokemonAreaResponse{}, err
		}
		return pokemonList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonAreaResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonAreaResponse{}, err
	}
	pokemonList := PokemonAreaResponse{}
	if err = json.Unmarshal(data, &pokemonList); err != nil {
		return PokemonAreaResponse{}, err
	}
	c.cache.Add(url, data)
	return pokemonList, nil
}
