package main

import (
	"time"

	"github.com/esh4d0w/bootdev-pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeapiClient,
	}
	startRepl(cfg)
}
