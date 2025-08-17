package main

import "fmt"

func commandExplore(c *config, args ...string) error {
	encounters, err := c.pokeApiClient.ListEncounters(args[0])
	if err != nil {
		return err
	}

	for _, encounter := range encounters.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
