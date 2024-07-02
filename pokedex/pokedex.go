package pokedex

import (
	"errors"
	"fmt"

	"github.com/VeryNotGood/Pokedex/internal/pokeapi"
)

type PokedexStorage struct {
	storage map[string]pokeapi.Pokemon
}

func NewEntry() PokedexStorage {
	p := PokedexStorage{
		storage: make(map[string]pokeapi.Pokemon),
	}
	return p
}

func (p *PokedexStorage) Add(key string, pokemon pokeapi.Pokemon) {
	p.storage[key] = pokemon
}

func (p *PokedexStorage) Get(key string) (pokeapi.Pokemon, bool) {
	entry, ok := p.storage[key]
	return entry, ok
}

func (p *PokedexStorage) List() error {
	if len(p.storage) == 0 {
		return errors.New("You have no pokemon!")
	}
	type pokedexList []string
	for n := range p.storage {
		fmt.Printf("  -%s\n", p.storage[n].Name)

	}
	return nil
}
