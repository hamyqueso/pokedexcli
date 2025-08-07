package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	// fmt.Println("help: Displays a help message")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap() error {

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area/", nil)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// fmt.Println(data)
	var location locationArea

	if err := json.Unmarshal(data, &location); err != nil {
		fmt.Println("Error unmarshalling")
		return fmt.Errorf("%w", err)
	}

	// fmt.Println(location)

	for _, area := range location.Results {
		fmt.Println(area.Name)
	}

	return nil

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

type locationArea struct {
	Next     string    `json:"next"`
	Previous any       `json:"previous"`
	Results  []results `json:"results"`
}

type results struct {
	Name string `json:"name"`
}

type config struct {
	Next     string
	Previous string
}

func cleanInput(text string) []string {
	var stringSlice []string
	for _, word := range strings.Fields(text) {
		stringSlice = append(stringSlice, strings.ToLower(word))
	}
	return stringSlice
}

func main() {
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
			commands[text].callback()
		} else {
			fmt.Println("Unknown command")
		}
	}

}
