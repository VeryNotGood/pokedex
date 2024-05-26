package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex help desk!")
	fmt.Println("Here are your help options:")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf("-%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
