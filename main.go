package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands map[string]cliCommand

func cleanInput(text string) []string{
	var stringSlice []string
	for _, word := range strings.Fields(text){
		stringSlice = append(stringSlice, strings.ToLower(word))
	}
	return stringSlice
}



func main() {
	commands = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a Help Message",
			callback: commandHelp,
		}, 
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: commandExit,
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
