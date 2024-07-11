package pokeapi

import (
	"net/http"
	"time"

	"github.com/Lanrey-waju/pokedexCLI/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheinterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheinterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
