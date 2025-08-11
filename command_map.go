package main

import (
	"fmt"
)

func commandMap(c *config) error { // fmt.Println(location)

	location, err := c.pokeApiClient.ListLocations(c.nextLocationsUrl)
	if err != nil {
		return err
	}

	c.nextLocationsUrl = location.Next
	c.previousLocationsUrl = location.Previous

	for _, area := range location.Results {
		fmt.Println(area.Name)
	}

	return nil

}

func commandMapB(c *config) error {
	if c.previousLocationsUrl == nil {
		fmt.Println("You're already on the first page")
	} else {

		location, err := c.pokeApiClient.ListLocations(c.previousLocationsUrl)
		if err != nil {
			return err
		}

		c.nextLocationsUrl = location.Next
		c.previousLocationsUrl = location.Previous

		for _, area := range location.Results {
			fmt.Println(area.Name)
		}
	}

	return nil
}
