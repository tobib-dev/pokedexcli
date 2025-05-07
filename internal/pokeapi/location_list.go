package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cacheData, exists := c.cache.Get(url)
	if exists {
		areas := LocationAreaResponse{}
		if err := json.Unmarshal(cacheData, &areas); err != nil {
			return LocationAreaResponse{}, err
		}
		return areas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	areas := LocationAreaResponse{}
	if err = json.Unmarshal(data, &areas); err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(url, data)
	return areas, nil
}
