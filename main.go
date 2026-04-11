package main

import (
	"time"

	"github.com/SamuelAboelkhir/pokedexcli/internal/pokeapi"
	"github.com/SamuelAboelkhir/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Minute)
	config := &Config{
		pokeapiClient: pokeClient,
		cache:         cache,
	}
	startRepl(config)
}
