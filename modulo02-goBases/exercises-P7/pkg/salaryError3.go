package main

import "fmt"

func newErrorTestFmt(salary int) {
	statusSalary := 15000
	if statusSalary <= 14999 {
		err := fmt.Errorf("minimo tributavel: %v", salary)
		fmt.Println("err: ", err)
	}
}
