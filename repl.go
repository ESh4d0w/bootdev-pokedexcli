package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/esh4d0w/bootdev-pokedexcli/internal/pokeapi"
	"github.com/esh4d0w/bootdev-pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokeCache     pokecache.Cache
	pokeDex       map[string]pokeapi.Pokemon
	nextMap       *string
	prevMap       *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex>")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	return strings.Fields(strings.ToLower(text))
}
