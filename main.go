package main

import (
	"bufio"
	"fmt"
	"github.com/hamyqueso/pokedexcli/internal/pokeapi"
	"os"
	"strings"
	"time"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	// fmt.Println("help: Displays a help message")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands map[string]cliCommand

type config struct {
	pokeApiClient        pokeapi.Client
	nextLocationsUrl     *string
	previousLocationsUrl *string
}

func cleanInput(text string) []string {
	var stringSlice []string
	for _, word := range strings.Fields(text) {
		stringSlice = append(stringSlice, strings.ToLower(word))
	}
	return stringSlice
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := config{
		pokeApiClient:        pokeClient,
		nextLocationsUrl:     nil,
		previousLocationsUrl: nil,
	}

	commands = map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays map locations. Calling it again displays the next 20 locations",
			callback:    commandMap,
		},
		"help": {
			name:        "help",
			description: "Displays a Help Message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		if _, exists := commands[text]; exists {
			commands[text].callback(&c)
		} else {
			fmt.Println("Unknown command")
		}
	}

}
