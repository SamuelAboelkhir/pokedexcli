package pokeapi

import (
	"net/http"
	"time"

	"github.com/SamuelAboelkhir/pokedexcli/internal/pokecache"
)

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
