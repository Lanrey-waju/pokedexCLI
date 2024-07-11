package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("pokedex takes no argument")
	}
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}
	fmt.Println("Your Pokedex: ")
	for k, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}
