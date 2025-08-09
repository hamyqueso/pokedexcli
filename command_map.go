package main

import (
	"fmt"
)

func commandMap(c *config) error { // fmt.Println(location)

	response, err := c.pokeApiClient.ListLocations(c.nextLocationsUrl)

	for _, area := range location.Results {
		fmt.Println(area.Name)
	}

	return nil

}
