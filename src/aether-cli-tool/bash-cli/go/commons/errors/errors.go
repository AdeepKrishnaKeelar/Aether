package commons

import (
	"log"
	"os"
)

// Function to produce Empty Value for key error.
func CommonError_EmptyValue(key string) {
	log.Printf("Empty value detected, expected for %s to have a value.", key)
	os.Exit(120)
}
