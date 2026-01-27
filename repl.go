package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "null",
	}

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := commandRegister()[commandName]
		if ok {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Printf("Trouble: %s\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandRegister() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandNextMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandPreviousMap,
		},
	}
	return commands
}
