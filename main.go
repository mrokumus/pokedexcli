package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	offset := 0
	for {
		fmt.Print("Pokedex > ")
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
		case "map":
			data, err := Map(offset)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, item := range data {
				fmt.Println(item)
			}
			offset += 20
		case "mapb":
			offset -= 20
			if offset < 0 {
				offset = 0
			}
			data, err := Map(offset)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, item := range data {
				fmt.Println(item)
			}
		default:
			use(words[0])
			break
		}
	}
}
