package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func startRepl(){
	type cliCommand struct {
		name string
		description string
		callback func() error
	}

	commands := map[string]cliCommand{
		"help": {
			name: "ℹ help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "⏻ exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n",cmd.name, cmd.description)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleanString := cleanInput(scanner.Text())
		
		if len(cleanString) == 0 {
			continue
		}
		
		commandName := cleanString[0]	
		
		if command, ok := commands[commandName];!ok {
			fmt.Println("Unknown command")
		} else {
			command.callback()
		}
			
		}
}


func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye! ✌️")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("This is a help message! :)\n")
	return nil
}

func cleanInput(text string) []string{
	lowerString := strings.ToLower(text)
	result := strings.Fields(lowerString)
	return result
}