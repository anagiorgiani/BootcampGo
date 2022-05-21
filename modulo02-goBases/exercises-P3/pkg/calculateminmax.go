package main

import (
	"errors"
	"fmt"
	"sort"
)


const (
	min = "minimum"
	ave = "average"
	max = "maximum"
)

func minValue(notas ...int) (float64, error) {
	sort.Ints(notas)

	l := len(notas)
	if l == 0 {
		return 0, nil
	}

	min := notas[0]
	fmt.Println(min)
	return float64(min), nil
}

func maxValue(notas ...int) (float64, error) {
	sort.Ints(notas)

	l := len(notas)
	if l == 0 {
		return 0, nil
	}

	max := notas[l-1]
	fmt.Println(max)
	return float64(max), nil
}

func averageValue(notas ...int) (float64, error) {
	var result int

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("Não pode num negativo")
		}

		result += value
	}

	return float64(result) / float64(len(notas)), nil
}

func calculateValues(calc string) func(notas ...int) (float64, error) {

	switch calc {
	case "minimum":
		return minValue
	case "average":
		return averageValue
	case "maximum":
		return maxValue
	}
	return nil
}
