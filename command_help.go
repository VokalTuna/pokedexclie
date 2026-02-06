package main

import "fmt"

func commandHelp(c *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, val := range commandRegister() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	fmt.Println()
	return nil
}
