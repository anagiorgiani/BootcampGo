package main

import (
	"fmt"
)

type error interface {
	Error() string
}

type mycustomerError struct {
	status int
	msg    string
}

func (e *mycustomerError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func mycustomerErrorsTest(status int) (int, error) {
	if status < 15000 {
		return status, &mycustomerError{
			status: status,
			msg:    "erro salario baixo",
		}
	}
	return status, nil
}
