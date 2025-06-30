package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
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

func commandMapNext(cfg *config) error {
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
func commandMapPrev(cfg *config) error {
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
