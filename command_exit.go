package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config) error {
	fmt.Println("catch ya later!")
	os.Exit(0)
	return nil
}
