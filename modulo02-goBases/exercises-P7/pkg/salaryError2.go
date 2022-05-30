package main

import (
	"errors"
	"fmt"
)

func newErrorTest(status int) {
	statusSalary := 15000
	if statusSalary <= 14999 {
		fmt.Println(errors.New("salario baixo"))
		return
	}
	fmt.Println("pagar imposto")

}
