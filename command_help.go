package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex CLI")
	fmt.Println("Here are your available commands")
	fmt.Println((" - help"))
	fmt.Println((" - exit"))

	return nil
}
