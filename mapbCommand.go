package main

import "fmt"

func mapbCommand(config *Config) error {
	if config.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := config.pokeapiClient.GetLocations(config.prev)
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
