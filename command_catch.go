package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	pokemon, err := c.pokeAPIClient.FindPokemon(args[0])
	if err != nil {
		return err
	}

	maxExp := 700
	pokeballChance := rand.Intn(maxExp)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if pokeballChance < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		c.caughtPokemon[pokemon.Name] = pokemon
	}

	return nil
}
