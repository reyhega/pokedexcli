package main

import (
	"fmt"
	"bufio"
	"os"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanString := cleanInput(input)
			fmt.Printf("Your command was: %s\n", cleanString[0])
		}
	}
}