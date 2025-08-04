package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string{
	var stringSlice []string
	for _, word := range strings.Fields(text){
		stringSlice = append(stringSlice, strings.ToLower(word))
	}
	return stringSlice
}

func main() {
	fmt.Println(cleanInput("  Hello World  "))
	fmt.Println(strings.TrimSpace("  hello  "))
}
