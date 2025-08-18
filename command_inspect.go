package main

import "fmt"

func commandInspect(c *config, args ...string) error {
	if data, exists := c.caughtPokemon[args[0]]; exists {
		fmt.Printf("Name: %s\n", data.Name)
		fmt.Printf("Height: %d\n", data.Height)
		fmt.Printf("Weight: %d\n", data.Weight)
		fmt.Println("Stats:")
		for _, stat := range data.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemontype := range data.Types {
			fmt.Printf("  - %s\n", pokemontype.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
