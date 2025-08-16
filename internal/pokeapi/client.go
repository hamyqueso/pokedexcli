package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/hamyqueso/pokedexcli/internal/cache"
)

type Client struct {
	httpClient http.Client
	pokecache  pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		pokecache:  pokecache.NewCache(10 * time.Second),
	}
}
