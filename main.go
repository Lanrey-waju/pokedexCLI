package main

import (
	"time"

	"github.com/Lanrey-waju/pokedexCLI/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Minute*5),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
