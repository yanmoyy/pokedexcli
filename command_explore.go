package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong arguments count, need to be 1 (name of area)")
	}
	name := args[0]
	fmt.Printf("Exploring %s...\n", name)
	locationResp, err := cfg.pokeapiClient.LocationArea(name)
	if err != nil {
		return fmt.Errorf("LocationArea: %w", err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationResp.PokemonEncounters {
		pokemon := pokemonEncounter.Pokemon
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
