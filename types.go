package main

import (
	"github.com/SamuelAboelkhir/pokedexcli/internal/pokeapi"
)

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
