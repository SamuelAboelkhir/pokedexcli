package main

import "fmt"

func pokedexCommand(config *config, name *string) error {
	pokedex := config.pokedex
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
