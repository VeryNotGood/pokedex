package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	// if len(args) >= 0 {
	// return errors.New("Don't do that you silly goose")
	// }
	fmt.Println("Your Pokedex:")
	err := cfg.pokedex.List()
	if err != nil {
		return err
	}

	return nil
}
