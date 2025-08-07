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
	fmt.Println("Usage:\n")
	fmt.Println("help: Displays a help message")
	for _, command := range commands {
		if command.name != "help" {
			fmt.Printf("%s: %s\n", command.name, command.description)
		}
	}
	return nil
}

func commandMap() error {
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location/1/", nil)
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

	var location []location

	if err := json.Unmarshal(data, &location); err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println("this is the map function")
	fmt.Println(location)
	return nil

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

type location struct {
	Name string `json:"name"`
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
		"map": {
			name:        "map",
			description: "Displays map locations",
			callback:    commandMap,
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
