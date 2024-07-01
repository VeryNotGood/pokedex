package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const myExperience = 30
	catchDifficulty := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Trying to catch %s...\n", pokemon.Name)
	for i := 1; i <= int(catchDifficulty/10); i++ {
		fmt.Print(".")
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Printf(".\n")

	if catchDifficulty > myExperience {
		fmt.Printf("You did not catch %s!\n", pokemonName)
		return err
	}

	fmt.Printf("You caught %s!\n", pokemon.Name)

	return nil
}
