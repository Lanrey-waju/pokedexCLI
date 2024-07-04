package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("no previous page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
