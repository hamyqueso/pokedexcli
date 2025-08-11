package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(locationsURL *string) (LocationsResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	// url := "https://google.com"
	//
	if locationsURL != nil {
		url = *locationsURL
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

	var location LocationsResponse

	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationsResponse{}, errors.New("error unmarshalling")
	}

	return location, nil
}
