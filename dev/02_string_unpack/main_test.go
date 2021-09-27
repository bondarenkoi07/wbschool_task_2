package main

import (
	"fmt"
	"testing"
)

var testStrings = map[string]string{
	"a4bc2d5e": "aaaabccddddde",
	"a10b06":   "aaaaaaaaaabbbbbb",
	"abcd":     "abcd",
	"45":       "",
	"":         "",
}

func TestStringUnpack(t *testing.T) {
	for testString, output := range testStrings {
		fmt.Printf("test string: %s\n", testString)
		funcOutput := StringUnpack(testString)
		if !(funcOutput == output) {
			t.Fatalf("func output %s != prepared output %s", funcOutput, output)
		}
	}
}
