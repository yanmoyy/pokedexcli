package main

import (
	"fmt"

	"github.com/yanmoyy/pokedexcli/api"
)

const apiUrl = "https://pokeapi.co/api/v2/location-area"

func commandMap(cfg *config) error {
	url := cfg.Next
	if url == "" {
		url = apiUrl
	}
	batch, err := api.FetchLocationAreasBatch(url)
	if err != nil {
		return fmt.Errorf("FetchLocationAreasBatch: %w", err)
	}
	if batch.Next != "" {
		cfg.Next = batch.Next
	}
	if batch.Previous != "" {
		cfg.Previous = batch.Previous
	}
	for _, area := range batch.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	url := cfg.Previous
	if url == "" {
		url = apiUrl
	}
	batch, err := api.FetchLocationAreasBatch(url)
	if err != nil {
		return fmt.Errorf("FetchLocationAreasBatch: %w", err)
	}
	if batch.Next != "" {
		cfg.Next = batch.Next
	}
	if batch.Previous != "" {
		cfg.Previous = batch.Previous
	}
	for _, area := range batch.Results {
		fmt.Println(area.Name)
	}
	return nil
}
