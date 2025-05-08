package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonGet(pageURL *string) (PokemonResponse, error) {
	url := baseURL + "/pokemon"
	if pageURL != nil {
		url = *pageURL
	}

	pokemon := PokemonResponse{}
	cachedData, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonResponse{}, err
	}
	c.cache.Add(url, data)
	return pokemon, nil
}
