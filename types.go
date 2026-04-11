package main

import (
	"github.com/SamuelAboelkhir/pokedexcli/internal/pokeapi"
	"github.com/SamuelAboelkhir/pokedexcli/internal/pokecache"
)

type Config struct {
	pokeapiClient pokeapi.Client
	next          *string
	prev          *string
	cache         pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}
