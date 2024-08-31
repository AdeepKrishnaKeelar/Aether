package main

import (
	"fmt"
	"go-logic/model"
	"os"

	_ "github.com/joho/godotenv"
)

func main() {
	// If minimum number of args are not passed, then it is invalid.
	if len(os.Args) < 2 {
		err := model.CallError(100, "Command not passed error")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
