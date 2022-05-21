package main

import (
	"errors"
)

func average(notas ...int) (float64, error) {
	var result int

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("Não pode num negativo")
		}

		result += value
	}

	return float64(result) / float64(len(notas)), nil
}