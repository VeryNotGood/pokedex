package main

import "github.com/VeryNotGood/Pokedex/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	prevLocation  *string
	nextLocation  *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	replStart(&cfg)
}
