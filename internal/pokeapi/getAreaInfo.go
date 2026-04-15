package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetNames(name *string) (LocationPokemon, error) {
	url := baseUrl + locationAreaPath

	if name != nil {
		url = url + "/" + *name
	}

	if data, ok := c.cache.Get(url); ok {
		locationPokemon := LocationPokemon{}
		err := json.Unmarshal(data, &locationPokemon)
		if err != nil {
			return LocationPokemon{}, err
		}
		return locationPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationPokemon{}, err
	}

	locationPokemon := LocationPokemon{}

	if err := json.Unmarshal(data, &locationPokemon); err != nil {
		return LocationPokemon{}, err
	}

	c.cache.Add(url, data)
	return locationPokemon, nil
}
