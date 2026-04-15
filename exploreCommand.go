package main

import "fmt"

func exploreCommand(config *config, name *string) error {
	if name == nil {
		fmt.Println("You must provide the name of a location to be explored")
	}

	names, err := config.pokeapiClient.GetNames(name)
	if err != nil {
		return err
	}
	for _, name := range names.PokemonEncounters {
		fmt.Println(name.Pokemon.Name)
	}
	return nil
}
