package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to provide a pokemon name")
	}
	name := args[0]
	val, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Weight: %d\n", val.Weight)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Stats:\n")
	for _, stats := range val.Stats {
		fmt.Printf("-%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, types := range val.Types {
		fmt.Printf("-%s\n", types.Type.Name)
	}

	return nil
}
