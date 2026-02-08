package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to provide a pokemon name.")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if CatchValue > cfg.randomSource.Intn(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.caughtPokemon[pokemon.Name] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}
