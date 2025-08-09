package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListLocations(locationsURL *string) (LocationsResponse, error) {
	req, err := http.NewRequest("GET", *locationsURL, nil)
	if err != nil {
		return LocationsResponse{}, errors.New("error creating request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, errors.New("error getting response")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, errors.New("error reading response")
	}

	var locations LocationsResponse

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsResponse{}, errors.New("Error unmarshalling")
	}

	return locations, nil
}
