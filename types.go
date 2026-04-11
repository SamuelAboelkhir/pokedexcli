package main

import pokeapi "github.com/SamuelAboelkhir/pokedexcli/pokeAPI"

type Config struct {
	pokeapiClient pokeapi.Client
	next          *string
	prev          *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}
