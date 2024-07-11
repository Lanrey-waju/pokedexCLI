package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		cleaned := cleanInput(line)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help menu",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows 20 Pokemon location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "map",
			description: "Shows 20 Pokemon location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Prints the pokemoon in an area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempts to catch the pokemoon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Attempts to catch the pokemoon",
			callback:    callbackInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex CLI",
			callback:    commandExit,
		},
	}
}
