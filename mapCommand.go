package main

import (
	"encoding/json"
	"fmt"

	"github.com/SamuelAboelkhir/pokedexcli/internal/pokeapi"
)

func mapCommand(config *Config) error {
	locations := pokeapi.Locations{}
	var data []byte
	var ok bool
	if config.next != nil {
		data, ok = config.cache.Get(*config.next)
	}
	if ok {
		json.Unmarshal(data, &locations)
	} else {
		apiData, err := config.pokeapiClient.GetLocations(config.next)
		if err != nil {
			return err
		}
		data, err := json.Marshal(apiData)
		if err != nil {
			return err
		}
		if config.next != nil {
			config.cache.Add(*config.next, data)
		} else {
			config.cache.Add(pokeapi.BaseUrl, data)
		}
		locations = apiData
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.next = locations.Next
	config.prev = locations.Previous
	return nil
}
