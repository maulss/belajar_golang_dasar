package main

import "fmt"

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

type notFoundError struct {
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "maulana" {
		return &notFoundError{Message: "not found error"}
	}

	return nil
}

func main() {
	err := saveData("", nil)
	if err != nil {
		// terjadi error
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("validation error", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error", notFoundError.Error())
		} else {
			fmt.Println("error", err.Error())
		}
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
