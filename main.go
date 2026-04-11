package main

import (
	"time"

	pokeapi "github.com/SamuelAboelkhir/pokedexcli/pokeAPI"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
