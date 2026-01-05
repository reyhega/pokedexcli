package main

import "strings"

func cleanInput(text string) []string{
	lowerString := strings.ToLower(text)
	result := strings.Fields(lowerString)
	
	return result
}