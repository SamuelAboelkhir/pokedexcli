package main

import (
	"encoding/json"
	"fmt"

	"github.com/SamuelAboelkhir/pokedexcli/internal/pokeapi"
)

func mapbCommand(config *Config) error {
	if config.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations := pokeapi.Locations{}

	data, ok := config.cache.Get(*config.prev)
	if ok {
		json.Unmarshal(data, &locations)
	} else {
		apiData, err := config.pokeapiClient.GetLocations(config.prev)
		if err != nil {
			return err
		}
		data, err := json.Marshal(apiData)
		if err != nil {
			return err
		}
		config.cache.Add(*config.prev, data)
		locations = apiData
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.next = locations.Next
	config.prev = locations.Previous
	return nil
}
