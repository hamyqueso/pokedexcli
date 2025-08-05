package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func cleanInput(text string) []string{
	var stringSlice []string
	for _, word := range strings.Fields(text){
		stringSlice = append(stringSlice, strings.ToLower(word))
	}
	return stringSlice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		fmt.Printf("Your command was: %s\n", cleanText[0])
	}

}
