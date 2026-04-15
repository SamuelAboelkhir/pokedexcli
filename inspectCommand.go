package main

import (
	"errors"
	"fmt"
)

func inspectCommand(config *config, name *string) error {
	if name == nil {
		return errors.New("Please provide a pokemon name")
	}

	if _, ok := config.pokedex[*name]; ok == false {
		return errors.New("you have not caught that pokemon")
	}
	pokemon := config.pokedex[*name]
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, types := range pokemon.Types {
		fmt.Printf(" - %s\n", types.Type.Name)
	}
	return nil
}
