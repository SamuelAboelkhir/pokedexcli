package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchCommand(config *config, name *string) error {
	if name == nil {
		return errors.New("Pick a pokemon to try to catch")
	}

	pokemon, err := config.pokeapiClient.GetPokemon(name)

	if err != nil {
		return err
	}

	catchChance := rand.Intn(pokemon.BaseExperience * 2)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchChance > pokemon.BaseExperience {
		fmt.Println(pokemon.Name, "was caught!")
		config.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	return nil
}
