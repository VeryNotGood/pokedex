package main

import (
	"time"

	"github.com/VeryNotGood/Pokedex/internal/pokeapi"
	"github.com/VeryNotGood/Pokedex/pokedex"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       pokedex.PokedexStorage
	prevLocation  *string
	nextLocation  *string
}

func main() {
	p := pokedex.NewEntry()
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokedex:       p,
	}
	replStart(&cfg)
}
