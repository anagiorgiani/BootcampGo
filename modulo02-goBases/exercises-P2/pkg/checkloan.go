package main


import (
	"fmt"
)

func checkLoan(age uint8, activeTime uint8, salary float64) {
	if age <= 22 {
		fmt.Println("idade insuficiente")
		return
	}

	if activeTime <= 1 {
		fmt.Println("tempo de serviço insuficiente")
		return
	}
	fmt.Println("empréstimo concedido")

	if salary > 100000 {
		fmt.Println("Não tem juros")
	}
}