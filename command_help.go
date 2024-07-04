package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex CLI")
	fmt.Println("Here are your available commands")
	commands := getCommands()

	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
