package main

import "fmt"

func mapCommand(config *Config) error {
	locations, err := config.pokeapiClient.GetLocations(config.next)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.next = locations.Next
	config.prev = locations.Previous
	return nil
}
