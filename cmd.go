package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func exit() {
	fmt.Printf("Closing the Pokedex... Goodbye!")
}

func help() {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
}

func use(cmd string) {
	fmt.Println(cmd)
}

func Map(offset int) ([]string, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", offset)
	req, err := http.Get(url)
	if err != nil {
		return []string{}, fmt.Errorf("an error occurred %v", err)
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		return []string{}, fmt.Errorf("an error occurred %v", req.StatusCode)
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return []string{}, fmt.Errorf("an error occurred %v", err)
	}
	var locationAreaResponse LocationAreaResponse
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return []string{}, fmt.Errorf("an error occurred %v", err)
	}
	var locationAreas []string
	results := locationAreaResponse.Results
	for _, locationArea := range results {
		locationAreas = append(locationAreas, locationArea.Name)
	}
	return locationAreas, nil
}
