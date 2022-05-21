package main

const (
	catA = 3000
	catB = 1500
	catC = 1000
)

func categoryA(h int) float64 {
	var result float64 = catA * float64(h)
	if h > 160 {
		return float64(result) * 1.5
	} else {
		return result
	}
}

func categoryB(h int) float64 {
	var result float64 = catB * float64(h)
	if h > 160 {
		return float64(result) * 1.2
	} else {
		return result
	}
}

func categoryC(h int) float64 {
	var result float64 = catC * float64(h)
	return result
}

func calculateSalary(category string, h int) float64 {

	switch category {
	case "A":
		return categoryA(h)
	case "B":
		return categoryB(h)
	case "C":
		return categoryC(h)
	}
	return 0
}