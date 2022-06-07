package testing

import "fmt"

func divide(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("o denominador nao pode ser 0")
	}
	return num / den, nil
}
