package main

import (
	"errors"
	"fmt"
)

func commandNextMap(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Previous = locationResp.Previous
	cfg.Next = locationResp.Next

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandPreviousMap(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Previous = locationResp.Previous
	cfg.Next = locationResp.Next

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
