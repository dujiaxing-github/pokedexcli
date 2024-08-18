package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	avaiableCommands := getCommands()
	for _, cmd := range avaiableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
