package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	FieldA string
	FieldB string
	Err    error
}

func (ce CustomError) Error() string {
	errorInfo := fmt.Sprintf("FieldA:%s, FieldB:%s, Err value:%s", ce.FieldA, ce.FieldB, ce.Err.Error())
	return errorInfo
}

func main() {
	err := &CustomError{
		FieldA: "err info a",
		FieldB: "err info b",
		Err:    errors.New("test custom err"),
	}

	fmt.Println(err.Error())
}
