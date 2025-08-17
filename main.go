package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hamyqueso/pokedexcli/internal/pokeapi"
)

func commandExit(c *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args ...string) error {
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
	callback    func(*config, ...string) error
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
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "displays the possible pokemon encounters at a specified area",
			callback:    commandExplore,
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
		text := cleanInput(scanner.Text())
		if len(text) == 1 {
			if text[0] == "explore" {
				fmt.Println("The explore function requires a location argument")
			} else if _, exists := commands[text[0]]; exists {
				commands[text[0]].callback(&c)
			} else {
				fmt.Println("Unknown command")
			}
		} else {
			if text[0] == "explore" {
				commands["explore"].callback(&c, text[1])
			} else if _, exists := commands[text[0]]; exists {
				fmt.Printf("The %s command does not take a second argument\n", text[0])
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
