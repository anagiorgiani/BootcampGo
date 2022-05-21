package main

func discountSalary(salary float64) float64 {
	tax := 0.0
	if salary <= 50000 {
		tax = salary * 1.17

	}
	if salary > 50000 {
		tax = salary * 1.10
	}
	return tax
}
