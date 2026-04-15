package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := getCommands()[words[0]]
		var name *string
		if len(words) > 1 {
			name = &words[1]
		}

		if ok {
			err := command.callback(config, name)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Print("Unknown command\n")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Shows the first 20 locations were pokemon are located. Fetches the next 20 when recalled",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 locations were pokemon are located. Continues to go back 20 locations at a time when recalled",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore",
			description: "Shows all the pokemon that are available in an area. Must provide area name",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon. Must provide pokemon name",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a pokemon and returns some information about him. Must provide a pokemon name",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displayes all the caught pokemon",
			callback:    pokedexCommand,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
