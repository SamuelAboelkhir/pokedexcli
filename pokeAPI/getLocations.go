package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocations(URL *string) (Locations, error) {
	url := baseUrl

	if URL != nil {
		url = *URL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	locations := Locations{}

	if err := decoder.Decode(&locations); err != nil {
		return Locations{}, err
	}

	return locations, nil
}
