package main

import (
	"fmt"
)

func commandHelp(_ *commandConfig) error {
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}
