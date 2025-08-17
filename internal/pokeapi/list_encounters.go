package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(locationName string) (EncountersResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName

	var encounters EncountersResponse

	if data, exists := c.pokecache.Get(url); exists {
		err := json.Unmarshal(data, &encounters)
		if err != nil {
			return EncountersResponse{}, errors.New("error unmarshhalling")
		}
		// fmt.Println("accessed cache")
		return encounters, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return EncountersResponse{}, errors.New("error creating request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return EncountersResponse{}, errors.New("error getting response")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return EncountersResponse{}, errors.New("error reading response")
	}

	if err = json.Unmarshal(data, &encounters); err != nil {
		return EncountersResponse{}, errors.New("error unmarshhalling")
	}

	c.pokecache.Add(url, data)

	return encounters, nil
}
