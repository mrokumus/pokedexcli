package main

import "strings"

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	var words []string
	words = strings.Split(strings.TrimSpace(text), " ")
	return words
}
