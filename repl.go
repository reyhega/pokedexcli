package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"Users/reyhernandohernandez/Desktop/Bootdev/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleanString := cleanInput(scanner.Text())
		
		if len(cleanString) == 0 {
			continue
		}
		
		commandName := cleanString[0]

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
			
	}
	
		
}

func cleanInput(text string) []string{
	lowerString := strings.ToLower(text)
	result := strings.Fields(lowerString)
	return result
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "ℹ help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Get next page of locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Get previous page of locations",
			callback: commandMapb,
		},
		"exit": {
			name: "⏻ exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},

	}
}

