package main

import "errors"

func categorySalary(h int, value float64) (float64, error) {
	if h < 80 {
		err := errors.New("horas trabalhadas inferior a 80h")
		return 0, err
	}
	var result float64 = float64(h) * value
	if result > 20000.00 {
		return float64(result) * 0.9, nil
	} else {
		return result, nil
	}
}
