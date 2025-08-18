package main

import (
	"fmt"
)

func commandMap(c *config, args ...string) error { // fmt.Println(location)

	location, err := c.pokeAPIClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = location.Next
	c.previousLocationsURL = location.Previous

	for _, area := range location.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapB(c *config, args ...string) error {
	if c.previousLocationsURL == nil {
		fmt.Println("You're already on the first page")
	} else {

		location, err := c.pokeAPIClient.ListLocations(c.previousLocationsURL)
		if err != nil {
			return err
		}

		c.nextLocationsURL = location.Next
		c.previousLocationsURL = location.Previous

		for _, area := range location.Results {
			fmt.Println(area.Name)
		}
	}

	return nil
}
