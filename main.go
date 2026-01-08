package main

import (
	"time"
	"Users/reyhernandohernandez/Desktop/Bootdev/pokedexcli/internal/pokeapi"
)

func main () {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}