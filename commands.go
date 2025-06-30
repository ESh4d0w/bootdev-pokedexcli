package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Shows the [next] 20 Locations",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 Locations",
			callback:    commandMapPrev,
		},
		"explore": {
			name:        "explore <AreaName>",
			description: "Shows what pokemon can be found in the Area",
			callback:    commandExploreMap,
		},
		"catch": {
			name:        "catch <PokemonName>",
			description: "Tries to catch a Pokemon",
			callback:    commandCatchPoke,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMapNext(cfg *config, args ...string) error {
	laList, err := cfg.pokeapiClient.GetLocationAreaList(cfg.nextMap)
	if err != nil {
		return err
	}

	cfg.nextMap = laList.Next
	cfg.prevMap = laList.Previous

	for _, location := range laList.Results {
		fmt.Println(location.Name)
	}

	if laList.Next == nil {
		fmt.Println("\nThis was the last Page!")
		cfg.nextMap = nil
	}

	return nil
}
func commandMapPrev(cfg *config, args ...string) error {
	if cfg.prevMap == nil {
		fmt.Println("There is no previous page.\n Try using map!")
		return nil
	}
	laList, err := cfg.pokeapiClient.GetLocationAreaList(cfg.prevMap)
	if err != nil {
		return err
	}

	cfg.nextMap = laList.Next
	cfg.prevMap = laList.Previous

	for _, location := range laList.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExploreMap(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	areaName := args[0]
	la, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", la.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounters := range la.PokemonEncounters {
		fmt.Println(encounters.Pokemon.Name)
	}
	return nil
}

func commandCatchPoke(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonaname := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonaname)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	chance := rand.Intn(370)
	if chance > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokeDex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
