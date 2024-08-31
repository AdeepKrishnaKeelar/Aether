// All the common data structures used in the application will be stored here.
package model

import "fmt"

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
