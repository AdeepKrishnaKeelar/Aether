package main

import (
	"fmt"
	"go-logic/cmd"
	aetcli "go-logic/cmd/aetcli"
	"go-logic/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// If minimum number of args are not passed, then it is invalid.

	if len(os.Args) < 2 {
		err := model.CallError(model.CommandNotPassed, "Command not passed error")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Load the Env Variables.
	err := godotenv.Load()
	if err != nil {
		msg := err
		err = model.CallError(model.EnvVariablesNotLoaded, "Failed to load env variables!")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}

	switch os.Args[1] {
	case model.Discover_node:
		// Call the main logic for Discover Node.
		cmd.DiscoverNode(os.Args[2:])

	case model.AetCLI:
		switch os.Args[2] {
		case model.ListNode:
			fmt.Println("Listing the VMs available...")
			aetcli.ListNode()
		}

	default:
		msg := os.Args[1] + " command not recognised by system."
		err = model.CallError(model.CommandNotFound, "Command not recognised!")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}
}
