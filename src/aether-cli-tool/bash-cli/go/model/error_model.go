// All the common data structures used in the application will be stored here.
package model

import (
	"fmt"
	"os"
)

// Set error codes as a const.
const (
	CommandNotFound       = 1
	CommandNotPassed      = 2
	EnvVariablesNotLoaded = 3
	FlagsNotLoaded        = 4
	EmptyFlagPassed       = 5
)

type RequestError struct {
	ErrorCode int
	ErrorMsg  string
}

// Function to implement the Error() method to satisfy the interface
func (r *RequestError) Error() string {
	return fmt.Sprintf("Status -- %d, Error -- %s.", r.ErrorCode, r.ErrorMsg)
}

// Function to create a custom error.
func CallError(code int, msg string) *RequestError {
	return &RequestError{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

// Function to produce Empty Value for key error.
func CommonError_EmptyValue(key string) {
	msg := "Empty value detected, expected for " + key + " to have a value."
	err := CallError(EmptyFlagPassed, msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
