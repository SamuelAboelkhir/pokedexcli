package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(URL, name *string) (Locations, error) {
	url := baseUrl + locationAreaPath

	if name != nil {
		url = url + "/" + *name
	}

	if URL != nil {
		url = *URL
	}
	if data, ok := c.cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return Locations{}, err
		}
		return locations, nil
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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}

	if err := json.Unmarshal(data, &locations); err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, data)
	return locations, nil
}
