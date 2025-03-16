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
