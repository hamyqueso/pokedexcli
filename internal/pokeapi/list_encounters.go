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

	return encounters, nil
}
