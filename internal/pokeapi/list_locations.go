// Pacakge pokeapi handles the http requests to the pokemon api
// found at https://pokeapi.co.

package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(locationsURL *string) (LocationsResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	var location LocationsResponse

	if locationsURL != nil {
		url = *locationsURL
	}

	if data, ok := c.pokecache.Get(url); ok {
		err := json.Unmarshal(data, &location)
		if err != nil {
			return LocationsResponse{}, errors.New("error unmarshalling")
		}
		fmt.Println("accessed cache")
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResponse{}, errors.New("error creating request")
	}

	// fmt.Println("created request")

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return LocationsResponse{}, errors.New("error getting response")
	}

	// fmt.Printf("got response: %d", res.StatusCode)

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, errors.New("error reading response")
	}

	// fmt.Println("read data")

	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationsResponse{}, errors.New("error unmarshalling")
	}

	c.pokecache.Add(url, data)

	return location, nil
}
