package commons

import (
	"fmt"
	"os"
)

// Function to produce Empty Value for key error.
func CommonError_EmptyValue(key string) {
	fmt.Printf("Empty value detected, expected for %s to have a value.", key)
	os.Exit(1)
}
