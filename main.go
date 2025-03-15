package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Hello, World!\n")
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	var words []string
	words = strings.Split(strings.TrimSpace(text), " ")
	return words
}
