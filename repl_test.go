package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input string
		expected []string
	}{
		"Lead & Trail":
	{
		input: " hello world ",
		expected: []string{"hello", "world"},
	},
		"Trail space":
	{
		input: "hello world ",
		expected: []string{"hello", "world"},
	},
		"Double space":
	{
		input: "hello  world",
		expected: []string{"hello", "world"},
	},
		"Simple":
	{
		input: "hello world",
		expected: []string{"hello", "world"},
	},
	}
	for name, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected){
			t.Errorf("%s failed --> expecting: %v | actual: %v", name, c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s failed --> expecting: %v | actual: %v", name, expectedWord, word)
			}
		}
	}

}
