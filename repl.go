package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := getCommands()[words[0]]

		if ok {
			err := command.callback(config)
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
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
