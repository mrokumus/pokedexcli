package main

import "fmt"

func exit() {
	fmt.Printf("Closing the Pokedex... Goodbye!")
}

func help() {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
}

func use(cmd string) {
	fmt.Println(cmd)
}
