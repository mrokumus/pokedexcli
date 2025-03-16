package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		words := strings.Fields(strings.ToLower(strings.TrimSpace(scanner.Text())))
		if len(words) == 0 {
			help()
			continue
		}
		switch words[0] {
		case "help":
			help()
			break
		case "exit":
			exit()
			os.Exit(0)
		default:
			use(words[0])
			break
		}
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	var words []string
	words = strings.Split(strings.TrimSpace(text), " ")
	return words
}

func exit() {
	fmt.Printf("Closing the Pokedex... Goodbye!")
}

func help() {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
}

func use(cmd string) {
	fmt.Println(cmd)
}
