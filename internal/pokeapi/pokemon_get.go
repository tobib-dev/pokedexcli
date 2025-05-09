package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonGet(pokeName string) (PokemonTraitsResponse, error) {
	url := baseURL + "/pokemon/" + pokeName

	pokemon := PokemonTraitsResponse{}
	cachedData, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return PokemonTraitsResponse{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonTraitsResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonTraitsResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonTraitsResponse{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonTraitsResponse{}, err
	}
	c.cache.Add(url, data)
	return pokemon, nil
}
