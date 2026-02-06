package main

import (
	"math/rand"
	"time"

	"github.com/VokalTuna/pokedexclie/internal/pokeapi"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &config{
		pokeapiClient: pokeClient,
		randomSource:  r,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
