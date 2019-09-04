package main

import "fmt"

type HTTPError struct {
	Status int
	Reason string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%v: %v", e.Status, e.Reason)
}

func MakeRequest() error {
	return HTTPError{
		Status: 400,
		Reason: "Bad Request",
	}
}

func main() {
	err := MakeRequest()
	if err != nil {
		fmt.Println(err) // 400: Bad Request
	}
}
