package main

import (
	"errors"
	"fmt"
)

type RequestError struct {
	ErrorCode int
	Err       error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status -- %d, err -- %v", r.ErrorCode, r.Err)
}

func doReq() *RequestError {
	return &RequestError{
		ErrorCode: 502,
		Err:       errors.New("some other error i have no idea"),
	}
}

func main() {
	//err := doReq()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	fmt.Println("Hey it ran correctly")
}
