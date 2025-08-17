package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) FindPokemon(name string) (PokemonResponse, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	var pokemon PokemonResponse

	if err = json.Unmarshal(data, &pokemon); err != nil {
		return PokemonResponse{}, err
	}

	return pokemon, nil
}
