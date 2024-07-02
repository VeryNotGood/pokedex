package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.pokedex.Get(pokemonName)
	if !ok {
		return errors.New("You don't have that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for s := range len(pokemon.Stats) {
		fmt.Printf("  -%s: %d\n", pokemon.Stats[s].Stat.Name, pokemon.Stats[s].BaseStat)
	}
	fmt.Println("Types:")
	for t := range len(pokemon.Types) {
		fmt.Printf("  -%s\n", pokemon.Types[t].Type.Name)
	}

	return nil
}
