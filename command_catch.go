package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if isCaught(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}

// isCaught determines if a Pokémon is caught based on its baseExperience.
// Returns true if caught, false otherwise.
func isCaught(baseExperience int) bool {
	const maxBaseExperience = 300    // Max reasonable base experience (e.g., legendaries ~300)
	const baseCatchProbability = 0.9 // Base probability for low baseExperience Pokémon

	catchProbability := baseCatchProbability - (float64(baseExperience)/float64(maxBaseExperience))*(baseCatchProbability/2)

	catchProbability = max(catchProbability, 0.1)

	roll := rand.Float64()

	return roll < catchProbability
}
