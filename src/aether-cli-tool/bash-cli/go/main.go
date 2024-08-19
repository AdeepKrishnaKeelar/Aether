package main

import (
	"fmt"
	"go-logic/cmd"
	"go-logic/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// If no command is passed, then throw an error.
	if len(os.Args) < 2 {
		fmt.Println("command not passed, please pass the command.")
		os.Exit(1)
	}

	// Set the global env variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error in loading env variables -- %s", err)
		os.Exit(1)
	}

	// All commands are part of the const model.
	switch os.Args[1] {
	case model.Discover_node:
		cmd.DiscoverNode(os.Args[2:])

	// In the case of a command not recognised, then throw an error.
	default:
		fmt.Printf("%s command not recognised.", os.Args[1])
		os.Exit(1)
	}
}
