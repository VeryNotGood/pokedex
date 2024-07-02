package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart(cfg *config) {

	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		input.Scan()
		text := input.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Returns next 20 locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns previous 20 locations",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location name}",
			description: "Returns Pokemon in the area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "Tries to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon name}",
			description: "Inspects named pokemon in pokedex",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all pokemon in pokedex",
			callback:    callbackPokedex,
		},
	}
}
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
