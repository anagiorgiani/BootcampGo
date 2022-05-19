package main

import "fmt"

func employee() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])
	var count = 0
	for _, v := range employees {
		if v > 21 {
			count++
		}

	}
	fmt.Println("Maiores de 21 anos:", count)
	employees["Federico"] = 25
	delete(employees, "Brenda")
	//fmt.Println(employees)

}