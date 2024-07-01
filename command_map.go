package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocation)
	if err != nil {
		return err
	}

	fmt.Println("List of locations: ")
	for _, location := range resp.Results {
		fmt.Printf("-- %s\n", location.Name)
	}
	cfg.nextLocation = resp.Next
	cfg.prevLocation = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocation == nil {
		return errors.New("there are no previous entries to display")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocation)
	if err != nil {
		return err
	}

	fmt.Println("List of locations: ")
	for _, location := range resp.Results {
		fmt.Printf("-- %s\n", location.Name)
	}
	cfg.nextLocation = resp.Next
	cfg.prevLocation = resp.Previous
	return nil
}
